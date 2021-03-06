package configuration

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/fabric8-services/fabric8-jenkins-idler/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestConfig_GetDebugMode(t *testing.T) {
	os.Setenv(debugMode, "false")
	c, _ := New("")
	assert.Equal(t, c.GetDebugMode(), false, "Debug Mode should be false")
}

func TestConfig_GetProxyURL(t *testing.T) {
	want := "https://proxy.openshift.io"
	os.Setenv(proxyURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetProxyURL(), "Proxy URL don't match")
}

func TestConfig_GetTenantURL(t *testing.T) {
	want := "https://tenent.openshift.io"
	os.Setenv(tenantURL, want)
	c, _ := New("")
	assert.Equal(t, want, c.GetTenantURL(), "Tenant URL")
}

func TestConfig_GetToggleURL(t *testing.T) {
	want := "https://toggle.openshift.io"
	os.Setenv(toggleURL, want)
	c, _ := New("")
	assert.Equal(t, c.GetToggleURL(), want, "Toggle URL Mismatch")
}

func TestConfig_GetAuthURL(t *testing.T) {
	want := "https://auth.openshift.io"
	os.Setenv(authURL, want)
	c, _ := New("")
	assert.Equal(t, c.GetAuthURL(), want, "Auth URL Mismatch")
}

func TestConfig_GetServiceAccountID(t *testing.T) {
	want := "1234567"
	os.Setenv(serviceAccountID, want)
	c, _ := New("")
	assert.Equal(t, c.GetServiceAccountID(), want,
		"Service Account ID Mismatch")
}

func TestConfig_GetServiceAccountSecret(t *testing.T) {
	want := "secretSvcAcnt"
	os.Setenv(serviceAccountSecret, want)
	c, _ := New("")
	assert.Equal(t, c.GetServiceAccountSecret(), want,
		"Service Account Secret Mismatch")
}

func TestConfig_GetAuthTokenKey(t *testing.T) {
	want := "authtokenkey"
	os.Setenv(authTokenKey, want)
	c, _ := New("")
	assert.Equal(t, c.GetAuthTokenKey(), want, "Auth Token Key Mismatch")
}

func TestConfig_GetAuthGrantType(t *testing.T) {
	want := "client_credentials"
	c, _ := New("")
	assert.Equal(t, c.GetAuthGrantType(), want, "Auth Grant Type Mismatch")
}

func TestConfig_GetIdleAfter(t *testing.T) {
	want := defaultIdleAfter
	c, _ := New("")
	assert.Equal(t, c.GetIdleAfter(), want, "Default Idle After Not Set")
}

func TestConfig_GetIdleLongBuild(t *testing.T) {
	want := defaultIdleLongBuild
	c, _ := New("")
	assert.Equal(t, c.GetIdleLongBuild(), want, "Idle Long Build")
}

func TestConfig_GetMaxRetries(t *testing.T) {
	want := defaultMaxRetries
	c, _ := New("")
	assert.Equal(t, c.GetMaxRetries(), want, "Max Retries failed")
}

func TestConfig_GetMaxRetriesQuietInterval(t *testing.T) {
	want := defaultMaxRetriesQuietInterval
	c, _ := New("")
	assert.Equal(t, c.GetMaxRetriesQuietInterval(),
		want, "Get Max Retries Quiet Interval Mismatch")
}

func TestConfig_GetCheckInterval(t *testing.T) {
	want := defaultCheckInterval
	c, _ := New("")
	assert.Equal(t, c.GetCheckInterval(), want, "Check Interval Mismatch")
}

func TestConfig_GetFixedUuids(t *testing.T) {
	os.Setenv(fixedUuids, "uuid1,uuid2,uuid3")
	want := []string{"uuid1", "uuid2", "uuid3"}
	c, _ := New("")
	assert.Equal(t, c.GetFixedUuids(), want, "FixedUUids Mismatch")
}

func TestConfig_Verify(t *testing.T) {
	os.Clearenv()
	os.Setenv(authTokenKey, "tokenkey")
	os.Setenv(serviceAccountSecret, "secret")
	os.Setenv(serviceAccountID, "1234567")
	os.Setenv(authURL, "https://auth.openshift.io")
	os.Setenv(toggleURL, "https://toggle.openshift.io")
	os.Setenv(tenantURL, "https://tenent.openshift.io")
	os.Setenv(proxyURL, "https://proxy.openshift.io")

	want := util.MultiError{}

	c, _ := New("")
	assert.Equal(t, c.Verify(), want, "Config Verification Failed")
}

func TestNew(t *testing.T) {
	_, err := New("fileNotFound.yaml")
	assert.Error(t, err,
		"Error expected when file not found")
}

func TestConfig_String(t *testing.T) {
	os.Clearenv()
	os.Setenv(authTokenKey, "tokenkey")
	os.Setenv(serviceAccountSecret, "secret")
	os.Setenv(serviceAccountID, "1234567")
	os.Setenv(authURL, "https://auth.openshift.io")
	os.Setenv(toggleURL, "https://toggle.openshift.io")
	os.Setenv(tenantURL, "https://tenent.openshift.io")
	os.Setenv(proxyURL, "https://proxy.openshift.io")

	c, _ := New("")
	assert.True(t, strings.Contains(c.String(), "jc_idle_after:"+strconv.Itoa(defaultIdleAfter)), "IdlerAfter Config String doesn't match")
	assert.True(t, strings.Contains(c.String(),
		"jc_max_retries_quiet_interval:"+
			strconv.Itoa(defaultMaxRetriesQuietInterval)),
		"Max retries Config String doesn't match")
	assert.True(t, strings.Contains(c.String(), "jc_auth_token_key:***"),
		"Auth Token key isn't ***")
	assert.True(t, strings.Contains(c.String(),
		"jc_service_account_secret:***"),
		"Service Account Secret isn't ***")
}
