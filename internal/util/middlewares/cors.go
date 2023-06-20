package middlewareinternal

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSConfig ...
func CORSConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           600,
	})
}

// AllowHeaders
const (
	HeaderAccessControlAllowHeaders = "Access-Control-Allow-Headers"
	HeaderAppName                   = "AppName"
	HeaderOrigin                    = "Origin"
	HeaderContentLength             = "Content-Length"
	HeaderContentType               = "Content-Type"
	HeaderAuthorization             = "Authorization"
	HeaderAcceptLanguage            = "Accept-Language"
	ResponseType                    = "responseType"
	HeaderVersion                   = "version"
	HeaderApiKey                    = "Api-Key"
	HeaderOSName                    = "OS-NAME"
	HeaderOSVersion                 = "OS-VERSION"
	HeaderPlatform                  = "PLATFORM"
	HeaderDeviceType                = "DeviceType"
	HeaderDeviceTypeWeb             = "Device-Type"
	HeaderBrowserName               = "Browser-Name"
	HeaderBrowserEmail              = "Browser-Email"
	HeaderBrowserVersion            = "Browser-Version"
	HeaderAppVersion                = "App-Version"
	HeaderAppVersionCode            = "App-Version-Code"
	HeaderUserAgent                 = "User-Agent"
	HeaderFCMToken                  = "Fcm-Token"
	HeaderModel                     = "DeviceModel"
	HeaderMerchantModel             = "Model"
	HeaderDeviceManufacturer        = "DeviceManufacturer"
	HeaderDeviceID                  = "DeviceId"
	HeaderIP                        = "IP"
	HeaderSource                    = "Source"
	HeaderWebDeviceID               = "Device-ID"
	HeaderMerchantDeviceID          = "DeviceId"
	HeaderMobileModel               = "MOBILE-MODEL"
	HeaderMobileVendor              = "MOBILE-VENDOR"
	HeaderManufacturer              = "Manufacturer"
	HeaderIsMobile                  = "Is-Mobile"
)

var (
	ListHeaderAllow = []string{
		HeaderOrigin,
		HeaderContentLength,
		HeaderContentType,
		HeaderAuthorization,
		HeaderAcceptLanguage,
		ResponseType,
		HeaderVersion,
		HeaderApiKey,
		HeaderOSName,
		HeaderOSVersion,
		HeaderPlatform,
		HeaderDeviceType,
		HeaderDeviceTypeWeb,
		HeaderBrowserName,
		HeaderBrowserVersion,
		HeaderAppVersion,
		HeaderAppVersionCode,
		HeaderUserAgent,
		HeaderFCMToken,
		HeaderModel,
		HeaderMerchantModel,
		HeaderDeviceID,
		HeaderMerchantDeviceID,
		HeaderIP,
		HeaderSource,
		HeaderAccessControlAllowHeaders,
		HeaderAppName,
		HeaderMobileVendor,
		HeaderMobileModel,
		HeaderDeviceManufacturer,
		HeaderWebDeviceID,
		HeaderManufacturer,
		HeaderIsMobile,
	}
)
