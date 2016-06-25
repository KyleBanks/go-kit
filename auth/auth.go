// The auth package provides generic authentication functionality.
package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	authCache Cache // cache is used for storing authentication tokens.
)

// Authenticator defines an interface for an authenticate-able User.
type Authenticator interface {
	Identifier() string     // Identifier returns a unique reference to this user.
	HashedPassword() string // HashedPassword returns the user's password hash.
}

// Cache defines the interface required for the authentication cache.
type Cache interface {
	PutString(key string, value string) (interface{}, error)
	GetString(key string) (string, error)
	Delete(key string) error

	// TODO: This shouldnt have to be here, but since Global references
	// the auth.Cache interface, we need to add it for now so that the Cache
	// package can be used globally.
	// TODO: Refactor global to reference cache.Cache instead of auth.Cache
	PutMarshaled(key string, value interface{}) (interface{}, error)
	GetMarshaled(key string, v interface{}) error
}

// Sets the Cache to use for authentication tokens.
func SetCache(c Cache) {
	authCache = c
}

// HashPassword returns a hashed version of the plain-text password provided.
func HashPassword(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash[:]), nil
}

// Authenticate validates an Authenticator based on it's password hash and the plain-text
// password provided.
func Authenticate(a Authenticator, plainTextPassword string) (AuthenticationTokenPair, error) {
	err := bcrypt.CompareHashAndPassword([]byte(a.HashedPassword()), []byte(plainTextPassword))
	if err != nil {
		return AuthenticationTokenPair{}, err
	}

	// Generate and cache a new token pair for this session
	return generateAndStoreTokens(a)
}

// Refresh generates a new token pair for a given authenticator.
func Refresh(a Authenticator, refreshToken string) (AuthenticationTokenPair, error) {
	newTokens, err := generateAndStoreTokens(a)
	if err != nil {
		return AuthenticationTokenPair{}, err
	}

	// Clear the old tokens from the cache
	if err := clearCachedTokens(refreshToken); err != nil {
		return AuthenticationTokenPair{}, err
	}

	return newTokens, nil
}

// GetIdentifierForAccessToken returns a user's identifier, as returned by
// the Authenticator interface, if it exists in the cache.
//
// If the identifier does not exist, and empty string and error will be returned.
func GetIdentifierForAccessToken(a string) (string, error) {
	return authCache.GetString(getAccessTokenCacheKey(a))
}

// GetIdentifierForRefreshToken returns a user's identifier, as returned by
// the Authenticator interface, if it exists in the cahce.
//
// If the identifier does not exist, an empty string and error will be returned.
func GetIdentifierForRefreshToken(r string) (string, error) {
	return authCache.GetString(getRefreshTokenCacheKey(r))
}

// generateAndStoreTokens creates and caches a new AuthenticationTokenPair.
func generateAndStoreTokens(a Authenticator) (AuthenticationTokenPair, error) {
	t := GenerateToken()
	if err := cacheTokens(t, a); err != nil {
		return AuthenticationTokenPair{}, err
	}

	return t, nil
}

// cacheTokens stores an access token and refresh token pair for an authenticated User.
func cacheTokens(t AuthenticationTokenPair, a Authenticator) error {
	if _, err := authCache.PutString(getAccessTokenCacheKey(t.AccessToken), a.Identifier()); err != nil {
		return err
	}

	if _, err := authCache.PutString(getRefreshTokenCacheKey(t.RefreshToken), a.Identifier()); err != nil {
		return err
	}

	if _, err := authCache.PutString(getRefreshToAccessTokenCacheKey(t.RefreshToken), t.AccessToken); err != nil {
		return err
	}

	return nil
}

// getAccessTokenCacheKey returns the access token cache key.
func getAccessTokenCacheKey(accessToken string) string {
	return fmt.Sprintf("accessToken:%s", accessToken)
}

// getRefreshTokenCacheKey returns the refresh token cache key.
func getRefreshTokenCacheKey(refreshToken string) string {
	return fmt.Sprintf("refreshToken:%s", refreshToken)
}

// getRefreshToAccessTokenCacheKey returns the refresh -> access token cache key.
func getRefreshToAccessTokenCacheKey(refreshToken string) string {
	return fmt.Sprintf("refreshToAccessToken:%s", refreshToken)
}

// clearCachedTokens clears all tokens associated to a refresh token.
func clearCachedTokens(r string) error {
	if a, err := authCache.GetString(getRefreshToAccessTokenCacheKey(r)); err != nil {
		return err
	} else if err = authCache.Delete(getAccessTokenCacheKey(a)); err != nil {
		return err
	} else if err = authCache.Delete(getRefreshTokenCacheKey(r)); err != nil {
		return err
	} else if err = authCache.Delete(getRefreshToAccessTokenCacheKey(r)); err != nil {
		return err
	}

	return nil
}
