package credentials_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weaveworks/eksctl/pkg/credentials"
	"github.com/weaveworks/eksctl/pkg/credentials/fakes"
)

type stubProvider struct {
	creds   credentials.Value
	expired bool
	err     error
}

func (s *stubProvider) Retrieve() (credentials.Value, error) {
	s.expired = false
	s.creds.ProviderName = "stubProvider"
	return s.creds, s.err
}

func (s *stubProvider) IsExpired() bool {
	return s.expired
}

type stubProviderExpirer struct {
	stubProvider
	expiration time.Time
}

func (s *stubProviderExpirer) ExpiresAt() time.Time {
	return s.expiration
}

var _ = Describe("filecache", func() {
	Context("credential cache has being used", func() {
		var (
			tmp string
			err error
		)
		BeforeEach(func() {
			tmp, err = ioutil.TempDir("", "filecache")
			Expect(err).ToNot(HaveOccurred())
			_ = os.Setenv(EksctlCacheFilenameEnvName, filepath.Join(tmp, "credentials.yaml"))
		})
		AfterEach(func() {
			_ = os.RemoveAll(tmp)
		})
		It("will provide a working file based cache", func() {
			c := credentials.NewCredentials(&stubProviderExpirer{
				stubProvider: stubProvider{
					creds: credentials.Value{
						AccessKeyID:     "id",
						SecretAccessKey: "secret",
						SessionToken:    "token",
						ProviderName:    "stubProvider",
					},
				},
			})
			fakeClock := &fakes.FakeClock{}
			fakeClock.NowReturns(time.Date(1981, 1, 1, 1, 1, 1, 1, time.UTC))
			p, err := NewFileCacheProvider("profile", c, fakeClock)
			Expect(err).ToNot(HaveOccurred())
			value, err := p.Retrieve()
			Expect(err).ToNot(HaveOccurred())
			Expect(value.AccessKeyID).To(Equal("id"))
			Expect(value.SecretAccessKey).To(Equal("secret"))
			Expect(value.SessionToken).To(Equal("token"))
			Expect(p.IsExpired()).NotTo(BeTrue())
			content, err := ioutil.ReadFile(filepath.Join(tmp, "credentials.yaml"))
			Expect(err).ToNot(HaveOccurred())
			Expect(string(content)).To(Equal(`profiles:
  profile:
    credential:
      accesskeyid: id
      secretaccesskey: secret
      sessiontoken: token
      providername: stubProvider
    expiration: 0001-01-01T00:00:00Z
`))
		})
		When("the cache expires", func() {
			It("will ask to refresh it", func() {
				c := credentials.NewCredentials(&stubProviderExpirer{
					stubProvider: stubProvider{
						creds: credentials.Value{
							AccessKeyID:     "id",
							SecretAccessKey: "secret",
							SessionToken:    "token",
							ProviderName:    "stubProvider",
						},
					},
				})
				fakeClock := &fakes.FakeClock{}
				fakeClock.NowReturns(time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC))
				p, err := NewFileCacheProvider("profile", c, fakeClock)
				Expect(err).ToNot(HaveOccurred())
				Expect(p.IsExpired()).To(BeTrue())
			})
		})
		When("the cache file already exists", func() {
			It("will retrieve its content", func() {
				content := []byte(`profiles:
  profile:
    credential:
      accesskeyid: storedID
      secretaccesskey: storedSecret
      sessiontoken: storedToken
      providername: stubProvider
    expiration: 0001-01-01T00:00:00Z
`)
				err := ioutil.WriteFile(filepath.Join(tmp, "credentials.yaml"), content, 0700)
				Expect(err).ToNot(HaveOccurred())
				c := credentials.NewCredentials(&stubProviderExpirer{})
				fakeClock := &fakes.FakeClock{}
				p, err := NewFileCacheProvider("profile", c, fakeClock)
				Expect(err).ToNot(HaveOccurred())
				creds, err := p.Retrieve()
				Expect(err).ToNot(HaveOccurred())
				Expect(creds.AccessKeyID).To(Equal("storedID"))
				Expect(creds.SecretAccessKey).To(Equal("storedSecret"))
				Expect(creds.SessionToken).To(Equal("storedToken"))
			})
		})
		When("no underlying credentials have been supplied", func() {
			It("returns an appropriate error", func() {
				fakeClock := &fakes.FakeClock{}
				_, err := NewFileCacheProvider("profile", nil, fakeClock)
				Expect(err).To(MatchError("no underlying Credentials object provided"))
			})
		})
		When("the underlying credentials provider doesn't support caching", func() {
			It("won't create a cache file", func() {
				fakeClock := &fakes.FakeClock{}
				fakeClock.NowReturns(time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC))
				p, err := NewFileCacheProvider("profile", credentials.NewStaticCredentials("id", "secret", "token"), fakeClock)
				Expect(err).ToNot(HaveOccurred())
				_, err = p.Retrieve()
				Expect(err).ToNot(HaveOccurred())
				_, err = os.Stat(filepath.Join(tmp, "credentials.yaml"))
				Expect(os.IsNotExist(err)).To(BeTrue())
			})

		})
		When("the cache file's permission is too broad", func() {
			It("will refuse to use that file", func() {
				content := []byte(`test:`)
				err := ioutil.WriteFile(filepath.Join(tmp, "credentials.yaml"), content, 0777)
				Expect(err).ToNot(HaveOccurred())
				c := credentials.NewCredentials(&stubProviderExpirer{})
				fakeClock := &fakes.FakeClock{}
				_, err = NewFileCacheProvider("profile", c, fakeClock)
				Expect(err).To(MatchError(fmt.Sprintf("cache file %s is not private", filepath.Join(tmp, "credentials.yaml"))))
			})
		})
		When("the cache data has been corrupted", func() {
			It("will return an appropriate error", func() {
				content := []byte(`not valid yaml`)
				err := ioutil.WriteFile(filepath.Join(tmp, "credentials.yaml"), content, 0600)
				Expect(err).ToNot(HaveOccurred())
				c := credentials.NewCredentials(&stubProviderExpirer{})
				fakeClock := &fakes.FakeClock{}
				_, err = NewFileCacheProvider("profile", c, fakeClock)
				Expect(err).To(MatchError(ContainSubstring("unable to parse file")))
			})
		})
	})
})