import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Separator } from '@/components/ui/separator';

export default function PrivacyPage() {
  return (
    <div className="container mx-auto py-12">
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-4">
          <h1 className="text-4xl font-bold tracking-tight">
            Chính Sách Bảo Mật
          </h1>
          <p className="text-muted-foreground text-lg">
            Cập nhật lần cuối: {new Date().toLocaleDateString('vi-VN')}
          </p>
        </div>

        <Separator />

        {/* Introduction */}
        <Card>
          <CardHeader>
            <CardTitle>Giới Thiệu</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chào mừng bạn đến với Niên Sử Việt. Chúng tôi cam kết bảo vệ quyền
              riêng tư và thông tin cá nhân của bạn. Chính sách bảo mật này giải
              thích cách chúng tôi thu thập, sử dụng, tiết lộ và bảo vệ thông
              tin của bạn khi bạn sử dụng dịch vụ của chúng tôi.
            </p>
          </CardContent>
        </Card>

        {/* Information Collection */}
        <Card>
          <CardHeader>
            <CardTitle>1. Thông Tin Chúng Tôi Thu Thập</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <h3 className="font-semibold mb-2">1.1. Thông Tin Cá Nhân</h3>
              <p className="text-muted-foreground mb-2">
                Chúng tôi có thể thu thập các thông tin cá nhân sau:
              </p>
              <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
                <li>Họ và tên</li>
                <li>Địa chỉ email</li>
                <li>Số điện thoại</li>
                <li>Thông tin đăng nhập (tên người dùng, mật khẩu)</li>
                <li>Ảnh đại diện (nếu có)</li>
              </ul>
            </div>

            <div>
              <h3 className="font-semibold mb-2">
                1.2. Thông Tin Tự Động Thu Thập
              </h3>
              <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
                <li>Địa chỉ IP</li>
                <li>Loại trình duyệt và phiên bản</li>
                <li>Hệ điều hành</li>
                <li>Thời gian truy cập và trang web được xem</li>
                <li>Cookies và dữ liệu theo dõi tương tự</li>
              </ul>
            </div>

            <div>
              <h3 className="font-semibold mb-2">
                1.3. Thông Tin Từ Bên Thứ Ba
              </h3>
              <p className="text-muted-foreground">
                Nếu bạn đăng nhập thông qua các dịch vụ bên thứ ba (Google,
                Facebook, v.v.), chúng tôi có thể nhận được thông tin từ các nền
                tảng đó theo chính sách của họ.
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Information Usage */}
        <Card>
          <CardHeader>
            <CardTitle>2. Cách Chúng Tôi Sử Dụng Thông Tin</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi sử dụng thông tin thu thập được cho các mục đích sau:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Cung cấp, vận hành và duy trì dịch vụ của chúng tôi</li>
              <li>Cải thiện, cá nhân hóa và mở rộng dịch vụ</li>
              <li>Hiểu và phân tích cách bạn sử dụng dịch vụ</li>
              <li>Phát triển sản phẩm, dịch vụ và tính năng mới</li>
              <li>Giao tiếp với bạn về dịch vụ, cập nhật và khuyến mãi</li>
              <li>Gửi thông báo và email liên quan đến tài khoản</li>
              <li>
                Phát hiện và ngăn chặn gian lận và các hoạt động bất hợp pháp
              </li>
              <li>Tuân thủ các nghĩa vụ pháp lý</li>
            </ul>
          </CardContent>
        </Card>

        {/* Information Sharing */}
        <Card>
          <CardHeader>
            <CardTitle>3. Chia Sẻ Thông Tin</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi không bán hoặc cho thuê thông tin cá nhân của bạn. Chúng
              tôi chỉ chia sẻ thông tin trong các trường hợp sau:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>
                <strong>Nhà cung cấp dịch vụ:</strong> Với các đối tác hỗ trợ
                vận hành dịch vụ
              </li>
              <li>
                <strong>Yêu cầu pháp lý:</strong> Khi được yêu cầu bởi pháp luật
                hoặc cơ quan chức năng
              </li>
              <li>
                <strong>Bảo vệ quyền lợi:</strong> Để bảo vệ quyền, tài sản hoặc
                an toàn của chúng tôi và người dùng
              </li>
              <li>
                <strong>Sự đồng ý của bạn:</strong> Với sự cho phép rõ ràng từ
                bạn
              </li>
              <li>
                <strong>Chuyển giao kinh doanh:</strong> Trong trường hợp sáp
                nhập, mua bán hoặc chuyển nhượng tài sản
              </li>
            </ul>
          </CardContent>
        </Card>

        {/* Data Security */}
        <Card>
          <CardHeader>
            <CardTitle>4. Bảo Mật Dữ Liệu</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi thực hiện các biện pháp bảo mật kỹ thuật và tổ chức để
              bảo vệ thông tin của bạn:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Mã hóa dữ liệu truyền tải (SSL/TLS)</li>
              <li>Mã hóa mật khẩu và thông tin nhạy cảm</li>
              <li>Kiểm soát truy cập nghiêm ngặt</li>
              <li>Giám sát và đánh giá bảo mật thường xuyên</li>
              <li>Sao lưu dữ liệu định kỳ</li>
            </ul>
            <p className="text-muted-foreground mt-4">
              Tuy nhiên, không có phương thức truyền tải qua internet hoặc lưu
              trữ điện tử nào là an toàn 100%. Chúng tôi không thể đảm bảo tính
              bảo mật tuyệt đối.
            </p>
          </CardContent>
        </Card>

        {/* Cookies */}
        <Card>
          <CardHeader>
            <CardTitle>5. Cookies và Công Nghệ Theo Dõi</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi sử dụng cookies và các công nghệ tương tự để:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Duy trì phiên đăng nhập của bạn</li>
              <li>Ghi nhớ tùy chọn và cài đặt của bạn</li>
              <li>Phân tích lưu lượng truy cập và tương tác người dùng</li>
              <li>Cá nhân hóa nội dung và quảng cáo</li>
            </ul>
            <p className="text-muted-foreground mt-4">
              Bạn có thể cấu hình trình duyệt để từ chối cookies, nhưng điều này
              có thể ảnh hưởng đến trải nghiệm sử dụng dịch vụ.
            </p>
          </CardContent>
        </Card>

        {/* User Rights */}
        <Card>
          <CardHeader>
            <CardTitle>6. Quyền Của Người Dùng</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Bạn có các quyền sau đối với thông tin cá nhân:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>
                <strong>Quyền truy cập:</strong> Yêu cầu bản sao thông tin cá
                nhân của bạn
              </li>
              <li>
                <strong>Quyền chỉnh sửa:</strong> Cập nhật hoặc sửa chữa thông
                tin không chính xác
              </li>
              <li>
                <strong>Quyền xóa:</strong> Yêu cầu xóa thông tin cá nhân của
                bạn
              </li>
              <li>
                <strong>Quyền hạn chế:</strong> Yêu cầu hạn chế xử lý thông tin
              </li>
              <li>
                <strong>Quyền di chuyển dữ liệu:</strong> Nhận thông tin của bạn
                ở định dạng có thể đọc được
              </li>
              <li>
                <strong>Quyền phản đối:</strong> Phản đối việc xử lý thông tin
                cá nhân
              </li>
              <li>
                <strong>Quyền rút lại sự đồng ý:</strong> Rút lại sự đồng ý đã
                cung cấp trước đó
              </li>
            </ul>
            <p className="text-muted-foreground mt-4">
              Để thực hiện các quyền này, vui lòng liên hệ với chúng tôi qua
              thông tin bên dưới.
            </p>
          </CardContent>
        </Card>

        {/* Data Retention */}
        <Card>
          <CardHeader>
            <CardTitle>7. Lưu Trữ Dữ Liệu</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi lưu giữ thông tin cá nhân của bạn trong thời gian cần
              thiết để:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Cung cấp dịch vụ cho bạn</li>
              <li>Tuân thủ các nghĩa vụ pháp lý</li>
              <li>Giải quyết tranh chấp</li>
              <li>Thực thi thỏa thuận của chúng tôi</li>
            </ul>
            <p className="text-muted-foreground mt-4">
              Khi bạn xóa tài khoản, chúng tôi sẽ xóa hoặc ẩn danh hóa thông tin
              cá nhân của bạn, trừ khi cần giữ lại cho các mục đích pháp lý.
            </p>
          </CardContent>
        </Card>

        {/* Children's Privacy */}
        <Card>
          <CardHeader>
            <CardTitle>8. Quyền Riêng Tư Của Trẻ Em</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Dịch vụ của chúng tôi không dành cho người dùng dưới 13 tuổi.
              Chúng tôi không cố ý thu thập thông tin cá nhân từ trẻ em dưới 13
              tuổi. Nếu bạn là phụ huynh và phát hiện con bạn đã cung cấp thông
              tin cho chúng tôi, vui lòng liên hệ để chúng tôi có thể xóa thông
              tin đó.
            </p>
          </CardContent>
        </Card>

        {/* Third Party Links */}
        <Card>
          <CardHeader>
            <CardTitle>9. Liên Kết Bên Thứ Ba</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Dịch vụ của chúng tôi có thể chứa liên kết đến các trang web hoặc
              dịch vụ của bên thứ ba. Chúng tôi không chịu trách nhiệm về chính
              sách bảo mật hoặc nội dung của các trang web đó. Chúng tôi khuyến
              khích bạn đọc chính sách bảo mật của bất kỳ trang web nào bạn truy
              cập.
            </p>
          </CardContent>
        </Card>

        {/* International Transfer */}
        <Card>
          <CardHeader>
            <CardTitle>10. Chuyển Giao Dữ Liệu Quốc Tế</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Thông tin của bạn có thể được chuyển giao và lưu trữ trên các máy
              chủ ở các quốc gia khác nhau. Chúng tôi đảm bảo rằng dữ liệu của
              bạn được bảo vệ phù hợp với chính sách bảo mật này bất kể nơi dữ
              liệu được xử lý.
            </p>
          </CardContent>
        </Card>

        {/* Policy Changes */}
        <Card>
          <CardHeader>
            <CardTitle>11. Thay Đổi Chính Sách</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi có thể cập nhật chính sách bảo mật này theo thời gian.
              Chúng tôi sẽ thông báo cho bạn về bất kỳ thay đổi nào bằng cách
              đăng chính sách mới trên trang này và cập nhật ngày "Cập nhật lần
              cuối". Chúng tôi khuyến khích bạn xem lại chính sách này định kỳ.
            </p>
          </CardContent>
        </Card>

        {/* Contact */}
        <Card>
          <CardHeader>
            <CardTitle>12. Liên Hệ</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Nếu bạn có bất kỳ câu hỏi hoặc thắc mắc nào về chính sách bảo mật
              này, vui lòng liên hệ với chúng tôi qua:
            </p>
            <div className="space-y-2 text-muted-foreground">
              <p>
                <strong>Email:</strong> privacy@niensu.viet
              </p>
              <p>
                <strong>Địa chỉ:</strong> [Địa chỉ công ty]
              </p>
              <p>
                <strong>Điện thoại:</strong> [Số điện thoại]
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Footer Note */}
        <div className="bg-muted p-6 rounded-lg">
          <p className="text-sm text-muted-foreground text-center">
            Bằng việc sử dụng dịch vụ của chúng tôi, bạn đồng ý với các điều
            khoản trong chính sách bảo mật này.
          </p>
        </div>
      </div>
    </div>
  );
}
