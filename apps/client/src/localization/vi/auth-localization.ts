import { ADMIN_ERROR_CODES } from './admin-error-codes';
import { ANONYMOUS_ERROR_CODES } from './anonymous-error-codes';
import { API_KEY_ERROR_CODES } from './api-key-error-codes';
import { BASE_ERROR_CODES } from './base-error-codes';
import { CAPTCHA_ERROR_CODES } from './captcha-error-codes';
import { EMAIL_OTP_ERROR_CODES } from './email-otp-error-codes';
import { GENERIC_OAUTH_ERROR_CODES } from './generic-oauth-error-codes';
import { HAVEIBEENPWNED_ERROR_CODES } from './haveibeenpwned-error-codes';
import { MULTI_SESSION_ERROR_CODES } from './multi-session-error-codes';
import { ORGANIZATION_ERROR_CODES } from './organization-error-codes';
import { PASSKEY_ERROR_CODES } from './passkey-error-codes';
import { PHONE_NUMBER_ERROR_CODES } from './phone-number-error-codes';
import { STRIPE_ERROR_CODES } from './stripe-localization';
import { TEAM_ERROR_CODES } from './team-error-codes';
import { TWO_FACTOR_ERROR_CODES } from './two-factor-error-codes';
import { USERNAME_ERROR_CODES } from './username-error-codes';

export const authLocalization = {
  /** @default "Ứng dụng" */
  APP: 'Ứng dụng',

  /** @default "Tài khoản" */
  ACCOUNT: 'Tài khoản',

  /** @default "Các tài khoản" */
  ACCOUNTS: 'Các tài khoản',

  /** @default "Chuyển đổi giữa các tài khoản đã đăng nhập." */
  ACCOUNTS_DESCRIPTION: 'Chuyển đổi giữa các tài khoản đã đăng nhập.',

  /** @default "Đăng nhập vào tài khoản bổ sung." */
  ACCOUNTS_INSTRUCTIONS: 'Đăng nhập vào tài khoản bổ sung.',

  /** @default "Thêm tài khoản" */
  ADD_ACCOUNT: 'Thêm tài khoản',

  /** @default "Thêm Passkey" */
  ADD_PASSKEY: 'Thêm Passkey',

  /** @default "Đã có tài khoản?" */
  ALREADY_HAVE_AN_ACCOUNT: 'Đã có tài khoản?',

  /** @default "Ảnh đại diện" */
  AVATAR: 'Ảnh đại diện',

  /** @default "Nhấp vào ảnh đại diện để tải lên ảnh tùy chỉnh từ thiết bị của bạn." */
  AVATAR_DESCRIPTION:
    'Nhấp vào ảnh đại diện để tải lên ảnh tùy chỉnh từ thiết bị của bạn.',

  /** @default "Ảnh đại diện là tùy chọn nhưng rất được khuyến nghị." */
  AVATAR_INSTRUCTIONS: 'Ảnh đại diện là tùy chọn nhưng rất được khuyến nghị.',

  /** @default "Mã dự phòng là bắt buộc" */
  BACKUP_CODE_REQUIRED: 'Mã dự phòng là bắt buộc',

  /** @default "Mã dự phòng" */
  BACKUP_CODES: 'Mã dự phòng',

  /** @default "Lưu các mã dự phòng này ở nơi an toàn. Bạn có thể sử dụng chúng để truy cập tài khoản nếu mất phương thức xác thực hai yếu tố." */
  BACKUP_CODES_DESCRIPTION:
    'Lưu các mã dự phòng này ở nơi an toàn. Bạn có thể sử dụng chúng để truy cập tài khoản nếu mất phương thức xác thực hai yếu tố.',

  /** @default "Mã dự phòng" */
  BACKUP_CODE_PLACEHOLDER: 'Mã dự phòng',

  /** @default "Mã dự phòng" */
  BACKUP_CODE: 'Mã dự phòng',

  /** @default "Hủy" */
  CANCEL: 'Hủy',

  /** @default "Đổi mật khẩu" */
  CHANGE_PASSWORD: 'Đổi mật khẩu',

  /** @default "Nhập mật khẩu hiện tại và mật khẩu mới." */
  CHANGE_PASSWORD_DESCRIPTION: 'Nhập mật khẩu hiện tại và mật khẩu mới.',

  /** @default "Vui lòng sử dụng tối thiểu 8 ký tự." */
  CHANGE_PASSWORD_INSTRUCTIONS: 'Vui lòng sử dụng tối thiểu 8 ký tự.',

  /** @default "Mật khẩu của bạn đã được thay đổi." */
  CHANGE_PASSWORD_SUCCESS: 'Mật khẩu của bạn đã được thay đổi.',

  /** @default "Xác nhận mật khẩu" */
  CONFIRM_PASSWORD: 'Xác nhận mật khẩu',

  /** @default "Xác nhận mật khẩu" */
  CONFIRM_PASSWORD_PLACEHOLDER: 'Xác nhận mật khẩu',

  /** @default "Xác nhận mật khẩu là bắt buộc" */
  CONFIRM_PASSWORD_REQUIRED: 'Xác nhận mật khẩu là bắt buộc',

  /** @default "Tiếp tục với Authenticator" */
  CONTINUE_WITH_AUTHENTICATOR: 'Tiếp tục với Authenticator',

  /** @default "Đã sao chép vào clipboard" */
  COPIED_TO_CLIPBOARD: 'Đã sao chép vào clipboard',

  /** @default "Sao chép vào clipboard" */
  COPY_TO_CLIPBOARD: 'Sao chép vào clipboard',

  /** @default "Sao chép tất cả mã" */
  COPY_ALL_CODES: 'Sao chép tất cả mã',

  /** @default "Tiếp tục" */
  CONTINUE: 'Tiếp tục',

  /** @default "Mật khẩu hiện tại" */
  CURRENT_PASSWORD: 'Mật khẩu hiện tại',

  /** @default "Mật khẩu hiện tại" */
  CURRENT_PASSWORD_PLACEHOLDER: 'Mật khẩu hiện tại',

  /** @default "Phiên hiện tại" */
  CURRENT_SESSION: 'Phiên hiện tại',

  /** @default "Cập nhật" */
  UPDATE: 'Cập nhật',

  /** @default "Xóa" */
  DELETE: 'Xóa',

  /** @default "Xóa ảnh đại diện" */
  DELETE_AVATAR: 'Xóa ảnh đại diện',

  /** @default "Xóa tài khoản" */
  DELETE_ACCOUNT: 'Xóa tài khoản',

  /** @default "Xóa vĩnh viễn tài khoản của bạn và tất cả nội dung. Hành động này không thể hoàn tác, vui lòng cân nhắc kỹ." */
  DELETE_ACCOUNT_DESCRIPTION:
    'Xóa vĩnh viễn tài khoản của bạn và tất cả nội dung. Hành động này không thể hoàn tác, vui lòng cân nhắc kỹ.',

  /** @default "Vui lòng xác nhận xóa tài khoản của bạn. Hành động này không thể hoàn tác, vui lòng cân nhắc kỹ." */
  DELETE_ACCOUNT_INSTRUCTIONS:
    'Vui lòng xác nhận xóa tài khoản của bạn. Hành động này không thể hoàn tác, vui lòng cân nhắc kỹ.',

  /** @default "Vui lòng kiểm tra email để xác minh việc xóa tài khoản của bạn." */
  DELETE_ACCOUNT_VERIFY:
    'Vui lòng kiểm tra email để xác minh việc xóa tài khoản của bạn.',

  /** @default "Tài khoản của bạn đã được xóa." */
  DELETE_ACCOUNT_SUCCESS: 'Tài khoản của bạn đã được xóa.',

  /** @default "Tắt xác thực hai yếu tố" */
  DISABLE_TWO_FACTOR: 'Tắt xác thực hai yếu tố',

  /** @default "Chọn nhà cung cấp để đăng nhập vào tài khoản của bạn" */
  DISABLED_CREDENTIALS_DESCRIPTION:
    'Chọn nhà cung cấp để đăng nhập vào tài khoản của bạn',

  /** @default "Chưa có tài khoản?" */
  DONT_HAVE_AN_ACCOUNT: 'Chưa có tài khoản?',

  /** @default "Hoàn tất" */
  DONE: 'Hoàn tất',

  /** @default "Email" */
  EMAIL: 'Email',

  /** @default "Nhập địa chỉ email bạn muốn sử dụng để đăng nhập." */
  EMAIL_DESCRIPTION: 'Nhập địa chỉ email bạn muốn sử dụng để đăng nhập.',

  /** @default "Vui lòng nhập địa chỉ email hợp lệ." */
  EMAIL_INSTRUCTIONS: 'Vui lòng nhập địa chỉ email hợp lệ.',

  /** @default "Email giống nhau" */
  EMAIL_IS_THE_SAME: 'Email giống nhau',

  /** @default "m@example.com" */
  EMAIL_PLACEHOLDER: 'm@example.com',

  /** @default "Địa chỉ email là bắt buộc" */
  EMAIL_REQUIRED: 'Địa chỉ email là bắt buộc',

  /** @default "Vui lòng kiểm tra email để xác minh thay đổi." */
  EMAIL_VERIFY_CHANGE: 'Vui lòng kiểm tra email để xác minh thay đổi.',

  /** @default "Vui lòng kiểm tra email để tìm liên kết xác minh." */
  EMAIL_VERIFICATION: 'Vui lòng kiểm tra email để tìm liên kết xác minh.',

  /** @default "Bật xác thực hai yếu tố" */
  ENABLE_TWO_FACTOR: 'Bật xác thực hai yếu tố',

  /** @default "không hợp lệ" */
  IS_INVALID: 'không hợp lệ',

  /** @default "là bắt buộc" */
  IS_REQUIRED: 'là bắt buộc',

  /** @default "giống nhau" */
  IS_THE_SAME: 'giống nhau',

  /** @default "Quên authenticator?" */
  FORGOT_AUTHENTICATOR: 'Quên authenticator?',

  /** @default "Quên mật khẩu" */
  FORGOT_PASSWORD: 'Quên mật khẩu',

  /** @default "Gửi liên kết đặt lại" */
  FORGOT_PASSWORD_ACTION: 'Gửi liên kết đặt lại',

  /** @default "Nhập email của bạn để đặt lại mật khẩu" */
  FORGOT_PASSWORD_DESCRIPTION: 'Nhập email của bạn để đặt lại mật khẩu',

  /** @default "Kiểm tra email của bạn để tìm liên kết đặt lại mật khẩu." */
  FORGOT_PASSWORD_EMAIL:
    'Kiểm tra email của bạn để tìm liên kết đặt lại mật khẩu.',

  /** @default "Quên mật khẩu?" */
  FORGOT_PASSWORD_LINK: 'Quên mật khẩu?',

  /** @default "Liên kết" */
  LINK: 'Liên kết',

  /** @default "Liên kết ma thuật" */
  MAGIC_LINK: 'Liên kết ma thuật',

  /** @default "Gửi liên kết ma thuật" */
  MAGIC_LINK_ACTION: 'Gửi liên kết ma thuật',

  /** @default "Nhập email của bạn để nhận liên kết ma thuật" */
  MAGIC_LINK_DESCRIPTION: 'Nhập email của bạn để nhận liên kết ma thuật',

  /** @default "Kiểm tra email của bạn để tìm liên kết ma thuật" */
  MAGIC_LINK_EMAIL: 'Kiểm tra email của bạn để tìm liên kết ma thuật',

  /** @default "Mã email" */
  EMAIL_OTP: 'Mã email',

  /** @default "Gửi mã" */
  EMAIL_OTP_SEND_ACTION: 'Gửi mã',

  /** @default "Xác minh mã" */
  EMAIL_OTP_VERIFY_ACTION: 'Xác minh mã',

  /** @default "Nhập email của bạn để nhận mã" */
  EMAIL_OTP_DESCRIPTION: 'Nhập email của bạn để nhận mã',

  /** @default "Vui lòng kiểm tra email của bạn để tìm mã xác minh." */
  EMAIL_OTP_VERIFICATION_SENT:
    'Vui lòng kiểm tra email của bạn để tìm mã xác minh.',

  /** @default "Tên" */
  NAME: 'Tên',

  /** @default "Vui lòng nhập tên đầy đủ hoặc tên hiển thị của bạn." */
  NAME_DESCRIPTION: 'Vui lòng nhập tên đầy đủ hoặc tên hiển thị của bạn.',

  /** @default "Vui lòng sử dụng tối đa 32 ký tự." */
  NAME_INSTRUCTIONS: 'Vui lòng sử dụng tối đa 32 ký tự.',

  /** @default "Tên" */
  NAME_PLACEHOLDER: 'Tên',

  /** @default "Mật khẩu mới" */
  NEW_PASSWORD: 'Mật khẩu mới',

  /** @default "Mật khẩu mới" */
  NEW_PASSWORD_PLACEHOLDER: 'Mật khẩu mới',

  /** @default "Mật khẩu mới là bắt buộc" */
  NEW_PASSWORD_REQUIRED: 'Mật khẩu mới là bắt buộc',

  /** @default "Mật khẩu một lần" */
  ONE_TIME_PASSWORD: 'Mật khẩu một lần',

  /** @default "Hoặc tiếp tục với" */
  OR_CONTINUE_WITH: 'Hoặc tiếp tục với',

  /** @default "Passkey" */
  PASSKEY: 'Passkey',

  /** @default "Passkeys" */
  PASSKEYS: 'Passkeys',

  /** @default "Quản lý passkeys của bạn để truy cập an toàn." */
  PASSKEYS_DESCRIPTION: 'Quản lý passkeys của bạn để truy cập an toàn.',

  /** @default "Truy cập tài khoản của bạn một cách an toàn mà không cần mật khẩu." */
  PASSKEYS_INSTRUCTIONS:
    'Truy cập tài khoản của bạn một cách an toàn mà không cần mật khẩu.',

  /** @default "Tài khoản cá nhân" */
  PERSONAL_ACCOUNT: 'Tài khoản cá nhân',

  /** @default "Khóa API" */
  API_KEYS: 'Khóa API',

  /** @default "Quản lý khóa API của bạn để truy cập an toàn." */
  API_KEYS_DESCRIPTION: 'Quản lý khóa API của bạn để truy cập an toàn.',

  /** @default "Tạo khóa API để truy cập tài khoản của bạn theo chương trình." */
  API_KEYS_INSTRUCTIONS:
    'Tạo khóa API để truy cập tài khoản của bạn theo chương trình.',

  /** @default "Tạo khóa API" */
  CREATE_API_KEY: 'Tạo khóa API',

  /** @default "Nhập tên duy nhất cho khóa API của bạn để phân biệt với các khóa khác." */
  CREATE_API_KEY_DESCRIPTION:
    'Nhập tên duy nhất cho khóa API của bạn để phân biệt với các khóa khác.',

  /** @default "Khóa API mới" */
  API_KEY_NAME_PLACEHOLDER: 'Khóa API mới',

  /** @default "Đã tạo khóa API" */
  API_KEY_CREATED: 'Đã tạo khóa API',

  /** @default "Vui lòng sao chép khóa API của bạn và lưu trữ ở nơi an toàn. Vì lý do bảo mật, chúng tôi không thể hiển thị lại." */
  CREATE_API_KEY_SUCCESS:
    'Vui lòng sao chép khóa API của bạn và lưu trữ ở nơi an toàn. Vì lý do bảo mật, chúng tôi không thể hiển thị lại.',

  /** @default "Không bao giờ hết hạn" */
  NEVER_EXPIRES: 'Không bao giờ hết hạn',

  /** @default "Hết hạn" */
  EXPIRES: 'Hết hạn',

  /** @default "Không có hạn" */
  NO_EXPIRATION: 'Không có hạn',

  /** @default "Tạo tổ chức" */
  CREATE_ORGANIZATION: 'Tạo tổ chức',

  /** @default "Tổ chức" */
  ORGANIZATION: 'Tổ chức',

  /** @default "Tên" */
  ORGANIZATION_NAME: 'Tên',

  /** @default "Công ty ABC" */
  ORGANIZATION_NAME_PLACEHOLDER: 'Công ty ABC',

  /** @default "Đây là tên hiển thị của tổ chức bạn." */
  ORGANIZATION_NAME_DESCRIPTION: 'Đây là tên hiển thị của tổ chức bạn.',

  /** @default "Vui lòng sử dụng tối đa 32 ký tự." */
  ORGANIZATION_NAME_INSTRUCTIONS: 'Vui lòng sử dụng tối đa 32 ký tự.',

  /** @default "URL Slug" */
  ORGANIZATION_SLUG: 'URL Slug',

  /** @default "Đây là không gian tên URL của tổ chức bạn." */
  ORGANIZATION_SLUG_DESCRIPTION: 'Đây là không gian tên URL của tổ chức bạn.',

  /** @default "Vui lòng sử dụng tối đa 48 ký tự." */
  ORGANIZATION_SLUG_INSTRUCTIONS: 'Vui lòng sử dụng tối đa 48 ký tự.',

  /** @default "cong-ty-abc" */
  ORGANIZATION_SLUG_PLACEHOLDER: 'cong-ty-abc',

  /** @default "Tạo tổ chức thành công" */
  CREATE_ORGANIZATION_SUCCESS: 'Tạo tổ chức thành công',

  /** @default "Mật khẩu" */
  PASSWORD: 'Mật khẩu',

  /** @default "Mật khẩu" */
  PASSWORD_PLACEHOLDER: 'Mật khẩu',

  /** @default "Mật khẩu là bắt buộc" */
  PASSWORD_REQUIRED: 'Mật khẩu là bắt buộc',

  /** @default "Mật khẩu không khớp" */
  PASSWORDS_DO_NOT_MATCH: 'Mật khẩu không khớp',

  /** @default "Nhà cung cấp" */
  PROVIDERS: 'Nhà cung cấp',

  /** @default "Kết nối tài khoản của bạn với dịch vụ bên thứ ba." */
  PROVIDERS_DESCRIPTION: 'Kết nối tài khoản của bạn với dịch vụ bên thứ ba.',

  /** @default "Khôi phục tài khoản" */
  RECOVER_ACCOUNT: 'Khôi phục tài khoản',

  /** @default "Khôi phục tài khoản" */
  RECOVER_ACCOUNT_ACTION: 'Khôi phục tài khoản',

  /** @default "Vui lòng nhập mã dự phòng để truy cập tài khoản của bạn" */
  RECOVER_ACCOUNT_DESCRIPTION:
    'Vui lòng nhập mã dự phòng để truy cập tài khoản của bạn',

  /** @default "Ghi nhớ tôi" */
  REMEMBER_ME: 'Ghi nhớ tôi',

  /** @default "Gửi lại mã" */
  RESEND_CODE: 'Gửi lại mã',

  /** @default "Gửi lại email xác minh" */
  RESEND_VERIFICATION_EMAIL: 'Gửi lại email xác minh',

  /** @default "Đặt lại mật khẩu" */
  RESET_PASSWORD: 'Đặt lại mật khẩu',

  /** @default "Lưu mật khẩu mới" */
  RESET_PASSWORD_ACTION: 'Lưu mật khẩu mới',

  /** @default "Nhập mật khẩu mới của bạn bên dưới" */
  RESET_PASSWORD_DESCRIPTION: 'Nhập mật khẩu mới của bạn bên dưới',

  /** @default "Đặt lại mật khẩu thành công" */
  RESET_PASSWORD_SUCCESS: 'Đặt lại mật khẩu thành công',

  /** @default "Yêu cầu thất bại" */
  REQUEST_FAILED: 'Yêu cầu thất bại',

  /** @default "Thu hồi" */
  REVOKE: 'Thu hồi',

  /** @default "Xóa khóa API" */
  DELETE_API_KEY: 'Xóa khóa API',

  /** @default "Bạn có chắc chắn muốn xóa khóa API này không?" */
  DELETE_API_KEY_CONFIRM: 'Bạn có chắc chắn muốn xóa khóa API này không?',

  /** @default "Khóa API" */
  API_KEY: 'Khóa API',

  /** @default "Đăng nhập" */
  SIGN_IN: 'Đăng nhập',

  /** @default "Đăng nhập" */
  SIGN_IN_ACTION: 'Đăng nhập',

  /** @default "Nhập email của bạn bên dưới để đăng nhập vào tài khoản" */
  SIGN_IN_DESCRIPTION: 'Nhập email của bạn bên dưới để đăng nhập vào tài khoản',

  /** @default "Nhập tên người dùng hoặc email để đăng nhập vào tài khoản" */
  SIGN_IN_USERNAME_DESCRIPTION:
    'Nhập tên người dùng hoặc email để đăng nhập vào tài khoản',

  /** @default "Đăng nhập với" */
  SIGN_IN_WITH: 'Đăng nhập với',

  /** @default "Đăng xuất" */
  SIGN_OUT: 'Đăng xuất',

  /** @default "Đăng ký" */
  SIGN_UP: 'Đăng ký',

  /** @default "Tạo tài khoản" */
  SIGN_UP_ACTION: 'Tạo tài khoản',

  /** @default "Nhập thông tin của bạn để tạo tài khoản" */
  SIGN_UP_DESCRIPTION: 'Nhập thông tin của bạn để tạo tài khoản',

  /** @default "Kiểm tra email của bạn để tìm liên kết xác minh." */
  SIGN_UP_EMAIL: 'Kiểm tra email của bạn để tìm liên kết xác minh.',

  /** @default "Phiên" */
  SESSIONS: 'Phiên',

  /** @default "Quản lý các phiên hoạt động và thu hồi quyền truy cập." */
  SESSIONS_DESCRIPTION:
    'Quản lý các phiên hoạt động và thu hồi quyền truy cập.',

  /** @default "Đặt mật khẩu" */
  SET_PASSWORD: 'Đặt mật khẩu',

  /** @default "Nhấp vào nút bên dưới để nhận email thiết lập mật khẩu cho tài khoản của bạn." */
  SET_PASSWORD_DESCRIPTION:
    'Nhấp vào nút bên dưới để nhận email thiết lập mật khẩu cho tài khoản của bạn.',

  /** @default "Cài đặt" */
  SETTINGS: 'Cài đặt',

  /** @default "Lưu" */
  SAVE: 'Lưu',

  /** @default "Bảo mật" */
  SECURITY: 'Bảo mật',

  /** @default "Chuyển đổi tài khoản" */
  SWITCH_ACCOUNT: 'Chuyển đổi tài khoản',

  /** @default "Tin tưởng thiết bị này" */
  TRUST_DEVICE: 'Tin tưởng thiết bị này',

  /** @default "Xác thực hai yếu tố" */
  TWO_FACTOR: 'Xác thực hai yếu tố',

  /** @default "Xác minh mã" */
  TWO_FACTOR_ACTION: 'Xác minh mã',

  /** @default "Vui lòng nhập mật khẩu một lần để tiếp tục" */
  TWO_FACTOR_DESCRIPTION: 'Vui lòng nhập mật khẩu một lần để tiếp tục',

  /** @default "Thêm lớp bảo mật bổ sung cho tài khoản của bạn." */
  TWO_FACTOR_CARD_DESCRIPTION:
    'Thêm lớp bảo mật bổ sung cho tài khoản của bạn.',

  /** @default "Vui lòng nhập mật khẩu của bạn để tắt 2FA." */
  TWO_FACTOR_DISABLE_INSTRUCTIONS: 'Vui lòng nhập mật khẩu của bạn để tắt 2FA.',

  /** @default "Vui lòng nhập mật khẩu của bạn để bật 2FA." */
  TWO_FACTOR_ENABLE_INSTRUCTIONS: 'Vui lòng nhập mật khẩu của bạn để bật 2FA.',

  /** @default "Đã bật xác thực hai yếu tố" */
  TWO_FACTOR_ENABLED: 'Đã bật xác thực hai yếu tố',

  /** @default "Đã tắt xác thực hai yếu tố" */
  TWO_FACTOR_DISABLED: 'Đã tắt xác thực hai yếu tố',

  /** @default "Xác thực hai yếu tố" */
  TWO_FACTOR_PROMPT: 'Xác thực hai yếu tố',

  /** @default "Quét mã QR bằng Authenticator của bạn" */
  TWO_FACTOR_TOTP_LABEL: 'Quét mã QR bằng Authenticator của bạn',

  /** @default "Gửi mã xác minh" */
  SEND_VERIFICATION_CODE: 'Gửi mã xác minh',

  /** @default "Hủy liên kết" */
  UNLINK: 'Hủy liên kết',

  /** @default "Cập nhật thành công" */
  UPDATED_SUCCESSFULLY: 'Cập nhật thành công',

  /** @default "Tên người dùng" */
  USERNAME: 'Tên người dùng',

  /** @default "Nhập tên người dùng bạn muốn sử dụng để đăng nhập." */
  USERNAME_DESCRIPTION: 'Nhập tên người dùng bạn muốn sử dụng để đăng nhập.',

  /** @default "Vui lòng sử dụng tối đa 32 ký tự." */
  USERNAME_INSTRUCTIONS: 'Vui lòng sử dụng tối đa 32 ký tự.',

  /** @default "Tên người dùng" */
  USERNAME_PLACEHOLDER: 'Tên người dùng',

  /** @default "(Tùy chọn)" */
  OPTIONAL_BRACKETS: '(Tùy chọn)',

  /** @default "Tên người dùng hoặc email" */
  SIGN_IN_USERNAME_PLACEHOLDER: 'Tên người dùng hoặc email',

  /** @default "Xác minh email của bạn" */
  VERIFY_YOUR_EMAIL: 'Xác minh email của bạn',

  /** @default "Vui lòng xác minh địa chỉ email của bạn. Kiểm tra hộp thư đến để tìm email xác minh. Nếu bạn chưa nhận được email, nhấp vào nút bên dưới để gửi lại." */
  VERIFY_YOUR_EMAIL_DESCRIPTION:
    'Vui lòng xác minh địa chỉ email của bạn. Kiểm tra hộp thư đến để tìm email xác minh. Nếu bạn chưa nhận được email, nhấp vào nút bên dưới để gửi lại.',

  /** @default "Quay lại" */
  GO_BACK: 'Quay lại',

  /** @default "Phiên của bạn đã hết hạn. Vui lòng đăng nhập lại." */
  SESSION_NOT_FRESH: 'Phiên của bạn đã hết hạn. Vui lòng đăng nhập lại.',

  /** @default "Tải lên ảnh đại diện" */
  UPLOAD_AVATAR: 'Tải lên ảnh đại diện',

  /** @default "Logo" */
  LOGO: 'Logo',

  /** @default "Nhấp vào logo để tải lên logo tùy chỉnh từ thiết bị của bạn." */
  LOGO_DESCRIPTION:
    'Nhấp vào logo để tải lên logo tùy chỉnh từ thiết bị của bạn.',

  /** @default "Logo là tùy chọn nhưng rất được khuyến nghị." */
  LOGO_INSTRUCTIONS: 'Logo là tùy chọn nhưng rất được khuyến nghị.',

  /** @default "Tải lên" */
  UPLOAD: 'Tải lên',

  /** @default "Tải lên logo" */
  UPLOAD_LOGO: 'Tải lên logo',

  /** @default "Xóa logo" */
  DELETE_LOGO: 'Xóa logo',

  /** @default "Chính sách bảo mật" */
  PRIVACY_POLICY: 'Chính sách bảo mật',

  /** @default "Điều khoản dịch vụ" */
  TERMS_OF_SERVICE: 'Điều khoản dịch vụ',

  /** @default "Trang này được bảo vệ bởi reCAPTCHA." */
  PROTECTED_BY_RECAPTCHA: 'Trang này được bảo vệ bởi reCAPTCHA.',

  /** @default "Bằng cách tiếp tục, bạn đồng ý với" */
  BY_CONTINUING_YOU_AGREE: 'Bằng cách tiếp tục, bạn đồng ý với',

  /** @default "Người dùng" */
  USER: 'Người dùng',

  /** @default "Tổ chức" */
  ORGANIZATIONS: 'Tổ chức',

  /** @default "Quản lý tổ chức và thành viên của bạn." */
  ORGANIZATIONS_DESCRIPTION: 'Quản lý tổ chức và thành viên của bạn.',

  /** @default "Tạo tổ chức để cộng tác với người dùng khác." */
  ORGANIZATIONS_INSTRUCTIONS: 'Tạo tổ chức để cộng tác với người dùng khác.',

  /** @default "Rời khỏi tổ chức" */
  LEAVE_ORGANIZATION: 'Rời khỏi tổ chức',

  /** @default "Bạn có chắc chắn muốn rời khỏi tổ chức này không?" */
  LEAVE_ORGANIZATION_CONFIRM:
    'Bạn có chắc chắn muốn rời khỏi tổ chức này không?',

  /** @default "Bạn đã rời khỏi tổ chức thành công." */
  LEAVE_ORGANIZATION_SUCCESS: 'Bạn đã rời khỏi tổ chức thành công.',

  /** @default "Quản lý tổ chức" */
  MANAGE_ORGANIZATION: 'Quản lý tổ chức',

  /** @default "Xóa thành viên" */
  REMOVE_MEMBER: 'Xóa thành viên',

  /** @default "Bạn có chắc chắn muốn xóa thành viên này khỏi tổ chức không?" */
  REMOVE_MEMBER_CONFIRM:
    'Bạn có chắc chắn muốn xóa thành viên này khỏi tổ chức không?',

  /** @default "Xóa thành viên thành công" */
  REMOVE_MEMBER_SUCCESS: 'Xóa thành viên thành công',

  /** @default "Mời thành viên" */
  INVITE_MEMBER: 'Mời thành viên',

  /** @default "Thành viên" */
  MEMBERS: 'Thành viên',

  /** @default "Thêm hoặc xóa thành viên và quản lý vai trò của họ." */
  MEMBERS_DESCRIPTION: 'Thêm hoặc xóa thành viên và quản lý vai trò của họ.',

  /** @default "Mời thành viên mới vào tổ chức của bạn." */
  MEMBERS_INSTRUCTIONS: 'Mời thành viên mới vào tổ chức của bạn.',

  /** @default "Gửi lời mời để thêm thành viên mới vào tổ chức của bạn." */
  INVITE_MEMBER_DESCRIPTION:
    'Gửi lời mời để thêm thành viên mới vào tổ chức của bạn.',

  /** @default "Vai trò" */
  ROLE: 'Vai trò',

  /** @default "Chọn vai trò" */
  SELECT_ROLE: 'Chọn vai trò',

  /** @default "Quản trị viên" */
  ADMIN: 'Quản trị viên',

  /** @default "Thành viên" */
  MEMBER: 'Thành viên',

  /** @default "Khách" */
  GUEST: 'Khách',

  /** @default "Chủ sở hữu" */
  OWNER: 'Chủ sở hữu',

  /** @default "Cập nhật vai trò cho thành viên này" */
  UPDATE_ROLE_DESCRIPTION: 'Cập nhật vai trò cho thành viên này',

  /** @default "Cập nhật vai trò" */
  UPDATE_ROLE: 'Cập nhật vai trò',

  /** @default "Cập nhật vai trò thành viên thành công" */
  MEMBER_ROLE_UPDATED: 'Cập nhật vai trò thành viên thành công',

  /** @default "Gửi lời mời" */
  SEND_INVITATION: 'Gửi lời mời',

  /** @default "Gửi lời mời thành công" */
  SEND_INVITATION_SUCCESS: 'Gửi lời mời thành công',

  /** @default "Lời mời đang chờ" */
  PENDING_INVITATIONS: 'Lời mời đang chờ',

  /** @default "Quản lý lời mời đang chờ cho tổ chức của bạn." */
  PENDING_INVITATIONS_DESCRIPTION:
    'Quản lý lời mời đang chờ cho tổ chức của bạn.',

  /** @default "Lời mời bạn đã nhận từ các tổ chức." */
  PENDING_USER_INVITATIONS_DESCRIPTION: 'Lời mời bạn đã nhận từ các tổ chức.',

  /** @default "Hủy lời mời" */
  CANCEL_INVITATION: 'Hủy lời mời',

  /** @default "Hủy lời mời thành công" */
  INVITATION_CANCELLED: 'Hủy lời mời thành công',

  /** @default "Chấp nhận lời mời" */
  ACCEPT_INVITATION: 'Chấp nhận lời mời',

  /** @default "Bạn đã được mời tham gia một tổ chức." */
  ACCEPT_INVITATION_DESCRIPTION: 'Bạn đã được mời tham gia một tổ chức.',

  /** @default "Chấp nhận lời mời thành công" */
  INVITATION_ACCEPTED: 'Chấp nhận lời mời thành công',

  /** @default "Từ chối lời mời thành công" */
  INVITATION_REJECTED: 'Từ chối lời mời thành công',

  /** @default "Chấp nhận" */
  ACCEPT: 'Chấp nhận',

  /** @default "Từ chối" */
  REJECT: 'Từ chối',

  /** @default "Lời mời này đã hết hạn" */
  INVITATION_EXPIRED: 'Lời mời này đã hết hạn',

  /** @default "Xóa tổ chức" */
  DELETE_ORGANIZATION: 'Xóa tổ chức',

  /** @default "Xóa vĩnh viễn tổ chức của bạn và tất cả nội dung. Hành động này không thể hoàn tác — vui lòng cân nhắc kỹ." */
  DELETE_ORGANIZATION_DESCRIPTION:
    'Xóa vĩnh viễn tổ chức của bạn và tất cả nội dung. Hành động này không thể hoàn tác — vui lòng cân nhắc kỹ.',

  /** @default "Xóa tổ chức thành công" */
  DELETE_ORGANIZATION_SUCCESS: 'Xóa tổ chức thành công',

  /** @default "Nhập slug tổ chức để tiếp tục:" */
  DELETE_ORGANIZATION_INSTRUCTIONS: 'Nhập slug tổ chức để tiếp tục:',

  /** @default "Slug tổ chức là bắt buộc" */
  SLUG_REQUIRED: 'Slug tổ chức là bắt buộc',

  /** @default "Slug không khớp" */
  SLUG_DOES_NOT_MATCH: 'Slug không khớp',

  // Teams
  /** @default "Nhóm" */
  TEAM: 'Nhóm',

  /** @default "Các nhóm" */
  TEAMS: 'Các nhóm',

  /** @default "Đang hoạt động" */
  TEAM_ACTIVE: 'Đang hoạt động',

  /** @default "Đặt làm hoạt động" */
  TEAM_SET_ACTIVE: 'Đặt làm hoạt động',

  /** @default "Tạo nhóm" */
  CREATE_TEAM: 'Tạo nhóm',

  /** @default "Tạo nhóm thành công" */
  CREATE_TEAM_SUCCESS: 'Tạo nhóm thành công',

  /** @default "Cập nhật nhóm" */
  UPDATE_TEAM: 'Cập nhật nhóm',

  /** @default "Cập nhật tên cho nhóm này" */
  UPDATE_TEAM_DESCRIPTION: 'Cập nhật tên cho nhóm này',

  /** @default "Bạn có chắc chắn muốn xóa nhóm này khỏi tổ chức không?" */
  REMOVE_TEAM_CONFIRM: 'Bạn có chắc chắn muốn xóa nhóm này khỏi tổ chức không?',

  /** @default "Thêm nhóm mới vào tổ chức của bạn." */
  CREATE_TEAM_INSTRUCTIONS: 'Thêm nhóm mới vào tổ chức của bạn.',

  /** @default "Tên nhóm" */
  TEAM_NAME: 'Tên nhóm',

  /** @default "Nhóm Kỹ thuật" */
  TEAM_NAME_PLACEHOLDER: 'Nhóm Kỹ thuật',

  /** @default "Đây là tên hiển thị của nhóm." */
  TEAM_NAME_DESCRIPTION: 'Đây là tên hiển thị của nhóm.',

  /** @default "Vui lòng sử dụng tối đa 64 ký tự." */
  TEAM_NAME_INSTRUCTIONS: 'Vui lòng sử dụng tối đa 64 ký tự.',

  /** @default "Quản lý các nhóm trong tổ chức của bạn." */
  TEAMS_DESCRIPTION: 'Quản lý các nhóm trong tổ chức của bạn.',

  /** @default "Bạn là thành viên của các nhóm sau." */
  USER_TEAMS_DESCRIPTION: 'Bạn là thành viên của các nhóm sau.',

  /** @default "Xóa nhóm" */
  DELETE_TEAM: 'Xóa nhóm',

  /** @default "Xóa vĩnh viễn nhóm này và tất cả nội dung." */
  DELETE_TEAM_DESCRIPTION: 'Xóa vĩnh viễn nhóm này và tất cả nội dung.',

  /** @default "Xóa nhóm thành công" */
  DELETE_TEAM_SUCCESS: 'Xóa nhóm thành công',

  /** @default "Nhập tên nhóm để tiếp tục:" */
  DELETE_TEAM_INSTRUCTIONS: 'Nhập tên nhóm để tiếp tục:',

  /** @default "Tên nhóm là bắt buộc" */
  TEAM_NAME_REQUIRED: 'Tên nhóm là bắt buộc',

  /** @default "Tên nhóm không khớp" */
  TEAM_NAME_DOES_NOT_MATCH: 'Tên nhóm không khớp',

  /** @default "Thành viên nhóm" */
  TEAM_MEMBERS: 'Thành viên nhóm',

  /** @default "Quản lý thành viên nhóm và vai trò của họ." */
  TEAM_MEMBERS_DESCRIPTION: 'Quản lý thành viên nhóm và vai trò của họ.',

  /** @default "Thêm thành viên nhóm" */
  ADD_TEAM_MEMBER: 'Thêm thành viên nhóm',

  /** @default "Xóa thành viên nhóm" */
  REMOVE_TEAM_MEMBER: 'Xóa thành viên nhóm',

  /** @default "Bạn có chắc chắn muốn xóa thành viên này khỏi nhóm không?" */
  REMOVE_TEAM_MEMBER_CONFIRM:
    'Bạn có chắc chắn muốn xóa thành viên này khỏi nhóm không?',

  /** @default "Xóa thành viên nhóm thành công" */
  REMOVE_TEAM_MEMBER_SUCCESS: 'Xóa thành viên nhóm thành công',

  /** @default "Thêm thành viên nhóm thành công" */
  ADD_TEAM_MEMBER_SUCCESS: 'Thêm thành viên nhóm thành công',

  /** @default "Cập nhật nhóm thành công" */
  UPDATE_TEAM_SUCCESS: 'Cập nhật nhóm thành công',

  /** @default "Quản lý thành viên nhóm" */
  MANAGE_TEAM_MEMBERS: 'Quản lý thành viên nhóm',

  /** @default "Tìm kiếm và thêm thành viên tổ chức vào nhóm này." */
  MANAGE_TEAM_MEMBERS_DESCRIPTION:
    'Tìm kiếm và thêm thành viên tổ chức vào nhóm này.',

  /** @default "Không tìm thấy nhóm" */
  NO_TEAMS_FOUND: 'Không tìm thấy nhóm',

  /** @default "thành viên" */
  MEMBER_SINGULAR: 'thành viên',

  /** @default "thành viên" */
  MEMBER_PLURAL: 'thành viên',

  /** @default "Không xác định" */
  UNKNOWN: 'Không xác định',

  ...BASE_ERROR_CODES,
  ...ADMIN_ERROR_CODES,
  ...ANONYMOUS_ERROR_CODES,
  ...API_KEY_ERROR_CODES,
  ...CAPTCHA_ERROR_CODES,
  ...EMAIL_OTP_ERROR_CODES,
  ...GENERIC_OAUTH_ERROR_CODES,
  ...HAVEIBEENPWNED_ERROR_CODES,
  ...MULTI_SESSION_ERROR_CODES,
  ...ORGANIZATION_ERROR_CODES,
  ...PASSKEY_ERROR_CODES,
  ...PHONE_NUMBER_ERROR_CODES,
  ...STRIPE_ERROR_CODES,
  ...TEAM_ERROR_CODES,
  ...TWO_FACTOR_ERROR_CODES,
  ...USERNAME_ERROR_CODES,
};

export type AuthLocalization = typeof authLocalization;
