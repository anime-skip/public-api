package config

// Database

func DatabaseURL() string {
	return requireEnvString("DATABASE_URL")
}
func DatabaseDisableSSL() bool {
	return envBool("DATABASE_DISABLE_SSL")
}
func DatabaseVersion() int {
	return envInt("DATABASE_VERSION")
}
func DatabaseEnableSeeding() bool {
	return envBool("DATABASE_ENABLE_SEEDING")
}

// Auth

func JWTSecret() string {
	return requireEnvString("JWT_SECRET")
}
func RecaptchaSecret() string {
	return requireEnvString("RECAPTCHA_SECRET")
}
func RecaptchaResponseAllowList() []string {
	return envStringArray("RECAPTCHA_RESPONSE_ALLOWLIST")
}

// Email Service

func EmailServiceHost() string {
	return requireEnvString("EMAIL_SERVICE_HOST")
}
func EmailServiceSecret() string {
	return requireEnvString("EMAIL_SERVICE_SECRET")
}
func EmailServiceEnabled() bool {
	return envBool("EMAIL_SERVICE_ENABLED")
}

// Better VRV

func BetterVRVAppID() string {
	return requireEnvString("BETTER_VRV_APP_ID")
}
func BetterVRVAPIKey() string {
	return requireEnvString("BETTER_VRV_API_KEY")
}

// Server

func EnableIntrospection() bool {
	return envBool("ENABLE_INTROSPECTION")
}
func EnablePlayground() bool {
	return envBool("ENABLE_PLAYGROUND")
}
func Port() int {
	return envInt("PORT")
}

// Logging

func LogLevel() int {
	return envIntOr("LOG_LEVEL", 0)
}
func DisableLogColors() bool {
	return envBool("DISABLE_LOG_COLORS")
}

// Discord

func DiscordBotToken() string {
	return envStringOr("DISCORD_BOT_TOKEN", "")
}
func DiscordAlertChannelID() string {
	return envString("DISCORD_ALERTS_CHANNEL_ID")
}
