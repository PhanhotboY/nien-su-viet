// These error codes are returned by the API
const EXTERNAL_ERROR_CODES = {
    VERIFICATION_FAILED: "Xác minh Captcha thất bại",
    MISSING_RESPONSE: "Thiếu phản hồi CAPTCHA",
    UNKNOWN_ERROR: "Đã xảy ra lỗi"
}

// These error codes are only visible in the server logs
const INTERNAL_ERROR_CODES = {
    MISSING_SECRET_KEY: "Thiếu khóa bí mật",
    SERVICE_UNAVAILABLE: "Dịch vụ CAPTCHA không khả dụng"
}

export const CAPTCHA_ERROR_CODES = {
    ...EXTERNAL_ERROR_CODES,
    ...INTERNAL_ERROR_CODES
}
