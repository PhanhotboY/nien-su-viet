export const API_KEY_ERROR_CODES = {
    INVALID_METADATA_TYPE: "metadata phải là một object hoặc undefined",
    REFILL_AMOUNT_AND_INTERVAL_REQUIRED:
        "refillAmount là bắt buộc khi refillInterval được cung cấp",
    REFILL_INTERVAL_AND_AMOUNT_REQUIRED:
        "refillInterval là bắt buộc khi refillAmount được cung cấp",
    USER_BANNED: "Người dùng đã bị cấm",
    UNAUTHORIZED_SESSION: "Phiên không được ủy quyền hoặc không hợp lệ",
    KEY_NOT_FOUND: "Không tìm thấy API Key",
    KEY_DISABLED: "API Key đã bị vô hiệu hóa",
    KEY_EXPIRED: "API Key đã hết hạn",
    USAGE_EXCEEDED: "API Key đã đạt đến giới hạn sử dụng",
    KEY_NOT_RECOVERABLE: "API Key không thể khôi phục",
    EXPIRES_IN_IS_TOO_SMALL:
        "expiresIn nhỏ hơn giá trị tối thiểu được xác định trước.",
    EXPIRES_IN_IS_TOO_LARGE:
        "expiresIn lớn hơn giá trị tối đa được xác định trước.",
    INVALID_REMAINING: "Số lượng còn lại quá lớn hoặc quá nhỏ.",
    INVALID_PREFIX_LENGTH:
        "Độ dài tiền tố quá lớn hoặc quá nhỏ.",
    INVALID_NAME_LENGTH: "Độ dài tên quá lớn hoặc quá nhỏ.",
    METADATA_DISABLED: "Metadata đã bị vô hiệu hóa.",
    RATE_LIMIT_EXCEEDED: "Vượt quá giới hạn tốc độ.",
    NO_VALUES_TO_UPDATE: "Không có giá trị nào để cập nhật.",
    KEY_DISABLED_EXPIRATION: "Giá trị hết hạn khóa tùy chỉnh đã bị vô hiệu hóa.",
    INVALID_API_KEY: "API key không hợp lệ.",
    INVALID_USER_ID_FROM_API_KEY: "ID người dùng từ API key không hợp lệ.",
    INVALID_API_KEY_GETTER_RETURN_TYPE:
        "API Key getter trả về loại khóa không hợp lệ. Mong đợi string.",
    SERVER_ONLY_PROPERTY:
        "Thuộc tính bạn đang cố gắng đặt chỉ có thể được đặt từ server auth instance."
}
