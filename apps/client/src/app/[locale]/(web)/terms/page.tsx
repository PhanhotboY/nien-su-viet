import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Separator } from '@/components/ui/separator';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { InfoIcon } from 'lucide-react';

export default function TermsPage() {
  return (
    <div className="container mx-auto py-12">
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-4">
          <h1 className="text-4xl font-bold tracking-tight">
            Điều Khoản Sử Dụng
          </h1>
          <p className="text-muted-foreground text-lg">
            Cập nhật lần cuối: {new Date().toLocaleDateString('vi-VN')}
          </p>
        </div>

        <Separator />

        {/* Alert */}
        <Alert>
          <InfoIcon className="h-4 w-4" />
          <AlertDescription>
            Vui lòng đọc kỹ các điều khoản này trước khi sử dụng dịch vụ. Việc
            sử dụng dịch vụ đồng nghĩa với việc bạn chấp nhận các điều khoản
            này.
          </AlertDescription>
        </Alert>

        {/* Introduction */}
        <Card>
          <CardHeader>
            <CardTitle>Giới Thiệu</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chào mừng bạn đến với Niên Sử Việt ("chúng tôi", "của chúng tôi"
              hoặc "dịch vụ"). Các điều khoản sử dụng này ("Điều Khoản") điều
              chỉnh việc truy cập và sử dụng trang web, ứng dụng và dịch vụ của
              chúng tôi. Bằng cách truy cập hoặc sử dụng dịch vụ, bạn đồng ý bị
              ràng buộc bởi các Điều Khoản này.
            </p>
          </CardContent>
        </Card>

        {/* Acceptance */}
        <Card>
          <CardHeader>
            <CardTitle>1. Chấp Nhận Điều Khoản</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Bằng cách truy cập và sử dụng dịch vụ của chúng tôi, bạn xác nhận
              rằng:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Bạn đã đọc và hiểu các Điều Khoản này</li>
              <li>Bạn đồng ý tuân thủ các Điều Khoản này</li>
              <li>Bạn đủ 13 tuổi trở lên</li>
              <li>
                Bạn có đầy đủ năng lực pháp lý để ký kết thỏa thuận ràng buộc
              </li>
              <li>
                Nếu bạn đại diện cho một tổ chức, bạn có thẩm quyền ràng buộc tổ
                chức đó
              </li>
            </ul>
            <p className="text-muted-foreground mt-4">
              Nếu bạn không đồng ý với bất kỳ phần nào của các Điều Khoản này,
              vui lòng không sử dụng dịch vụ.
            </p>
          </CardContent>
        </Card>

        {/* Account Registration */}
        <Card>
          <CardHeader>
            <CardTitle>2. Đăng Ký Tài Khoản</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <h3 className="font-semibold mb-2">2.1. Tạo Tài Khoản</h3>
              <p className="text-muted-foreground">
                Để sử dụng một số tính năng, bạn cần tạo tài khoản. Khi đăng ký,
                bạn phải:
              </p>
              <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4 mt-2">
                <li>Cung cấp thông tin chính xác, đầy đủ và cập nhật</li>
                <li>Bảo mật thông tin đăng nhập của bạn</li>
                <li>
                  Thông báo ngay cho chúng tôi về bất kỳ vi phạm bảo mật nào
                </li>
                <li>
                  Chịu trách nhiệm về mọi hoạt động xảy ra dưới tài khoản của
                  bạn
                </li>
              </ul>
            </div>

            <div>
              <h3 className="font-semibold mb-2">2.2. Hạn Chế Tài Khoản</h3>
              <p className="text-muted-foreground">Bạn không được:</p>
              <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4 mt-2">
                <li>Tạo nhiều tài khoản cho mục đích gian lận</li>
                <li>Chia sẻ tài khoản của bạn với người khác</li>
                <li>
                  Chuyển nhượng tài khoản mà không có sự đồng ý của chúng tôi
                </li>
                <li>Sử dụng tài khoản người khác mà không được phép</li>
              </ul>
            </div>
          </CardContent>
        </Card>

        {/* User Conduct */}
        <Card>
          <CardHeader>
            <CardTitle>3. Quy Tắc Ứng Xử</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Khi sử dụng dịch vụ, bạn đồng ý KHÔNG:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Vi phạm bất kỳ luật hoặc quy định nào</li>
              <li>
                Đăng nội dung bất hợp pháp, có hại, lừa đảo, quấy rối, hoặc xúc
                phạm
              </li>
              <li>Xâm phạm quyền sở hữu trí tuệ của người khác</li>
              <li>Phát tán virus, mã độc hoặc công nghệ có hại</li>
              <li>
                Thu thập thông tin cá nhân của người dùng khác mà không có sự
                đồng ý
              </li>
              <li>Can thiệp hoặc phá hoại dịch vụ</li>
              <li>Sử dụng bot, script hoặc công cụ tự động không được phép</li>
              <li>Mạo danh bất kỳ cá nhân hoặc tổ chức nào</li>
              <li>Spam hoặc gửi nội dung quảng cáo không mong muốn</li>
              <li>Khai thác lỗ hổng hoặc lạm dụng dịch vụ</li>
            </ul>
          </CardContent>
        </Card>

        {/* User Content */}
        <Card>
          <CardHeader>
            <CardTitle>4. Nội Dung Người Dùng</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <h3 className="font-semibold mb-2">4.1. Quyền Sở Hữu</h3>
              <p className="text-muted-foreground">
                Bạn giữ quyền sở hữu đối với nội dung bạn đăng tải. Tuy nhiên,
                bằng cách đăng nội dung, bạn cấp cho chúng tôi giấy phép toàn
                cầu, không độc quyền, miễn phí bản quyền, có thể chuyển nhượng
                để sử dụng, sao chép, phân phối, hiển thị và thực hiện nội dung
                đó liên quan đến dịch vụ.
              </p>
            </div>

            <div>
              <h3 className="font-semibold mb-2">4.2. Trách Nhiệm</h3>
              <p className="text-muted-foreground">
                Bạn chịu trách nhiệm hoàn toàn về nội dung bạn đăng tải. Bạn đảm
                bảo rằng:
              </p>
              <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4 mt-2">
                <li>Bạn có quyền đăng nội dung đó</li>
                <li>Nội dung không vi phạm quyền của bên thứ ba</li>
                <li>Nội dung tuân thủ các Điều Khoản này</li>
              </ul>
            </div>

            <div>
              <h3 className="font-semibold mb-2">4.3. Kiểm Duyệt</h3>
              <p className="text-muted-foreground">
                Chúng tôi có quyền (nhưng không có nghĩa vụ) xem xét, giám sát,
                chỉnh sửa hoặc xóa nội dung người dùng mà chúng tôi cho là vi
                phạm Điều Khoản này hoặc gây hại cho dịch vụ.
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Intellectual Property */}
        <Card>
          <CardHeader>
            <CardTitle>5. Quyền Sở Hữu Trí Tuệ</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Dịch vụ và nội dung gốc của nó (ngoại trừ nội dung người dùng),
              các tính năng và chức năng là và sẽ vẫn là tài sản độc quyền của
              Niên Sử Việt và các bên cấp phép của chúng tôi. Dịch vụ được bảo
              vệ bởi bản quyền, nhãn hiệu và các luật khác. Các nhãn hiệu và
              logo của chúng tôi không được sử dụng mà không có sự cho phép bằng
              văn bản trước.
            </p>
          </CardContent>
        </Card>

        {/* Third Party Services */}
        <Card>
          <CardHeader>
            <CardTitle>6. Dịch Vụ Bên Thứ Ba</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Dịch vụ của chúng tôi có thể chứa liên kết đến các trang web hoặc
              dịch vụ của bên thứ ba. Chúng tôi không kiểm soát và không chịu
              trách nhiệm về nội dung, chính sách bảo mật hoặc thực tiễn của bất
              kỳ trang web hoặc dịch vụ bên thứ ba nào. Bạn thừa nhận và đồng ý
              rằng chúng tôi không chịu trách nhiệm về bất kỳ thiệt hại nào do
              việc sử dụng các dịch vụ đó gây ra.
            </p>
          </CardContent>
        </Card>

        {/* Termination */}
        <Card>
          <CardHeader>
            <CardTitle>7. Chấm Dứt</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <h3 className="font-semibold mb-2">7.1. Bởi Bạn</h3>
              <p className="text-muted-foreground">
                Bạn có thể ngừng sử dụng dịch vụ và xóa tài khoản của mình bất
                kỳ lúc nào thông qua cài đặt tài khoản.
              </p>
            </div>

            <div>
              <h3 className="font-semibold mb-2">7.2. Bởi Chúng Tôi</h3>
              <p className="text-muted-foreground">
                Chúng tôi có quyền đình chỉ hoặc chấm dứt quyền truy cập của bạn
                vào dịch vụ ngay lập tức, không cần thông báo trước, vì bất kỳ
                lý do gì, bao gồm nhưng không giới hạn ở:
              </p>
              <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4 mt-2">
                <li>Vi phạm các Điều Khoản này</li>
                <li>Yêu cầu của cơ quan pháp luật hoặc chính phủ</li>
                <li>Hoạt động gian lận, bất hợp pháp hoặc không phù hợp</li>
                <li>Ngừng cung cấp dịch vụ</li>
              </ul>
            </div>

            <div>
              <h3 className="font-semibold mb-2">7.3. Hậu Quả</h3>
              <p className="text-muted-foreground">
                Khi chấm dứt, quyền sử dụng dịch vụ của bạn sẽ ngay lập tức chấm
                dứt. Chúng tôi có thể xóa hoặc ẩn danh hóa nội dung và dữ liệu
                của bạn.
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Disclaimer */}
        <Card>
          <CardHeader>
            <CardTitle>8. Tuyên Bố Từ Chối Trách Nhiệm</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              DỊCH VỤ ĐƯỢC CUNG CẤP TRÊN CƠ SỞ "NHƯ VẬY" VÀ "NHƯ CÓ SẴN" MÀ
              KHÔNG CÓ BẤT KỲ BẢO ĐẢM NÀO, DÙ RÕ RÀNG HAY NGỤ Ý. CHÚNG TÔI KHÔNG
              BẢO ĐẢM RẰNG DỊCH VỤ SẼ:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Hoạt động không bị gián đoạn hoặc an toàn</li>
              <li>Không có lỗi</li>
              <li>Đáp ứng yêu cầu của bạn</li>
              <li>Cung cấp kết quả chính xác hoặc đáng tin cậy</li>
            </ul>
            <p className="text-muted-foreground mt-4">
              BẠN SỬ DỤNG DỊCH VỤ HOÀN TOÀN TỰ CHỊU RỦI RO. CHÚNG TÔI KHÔNG CHỊU
              TRÁCH NHIỆM VỀ BẤT KỲ NỘI DUNG, MÃ HOẶC HÀNH VI CỦA BẤT KỲ BÊN THỨ
              BA NÀO.
            </p>
          </CardContent>
        </Card>

        {/* Limitation of Liability */}
        <Card>
          <CardHeader>
            <CardTitle>9. Giới Hạn Trách Nhiệm</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              TRONG MỌI TRƯỜNG HỢP, CHÚNG TÔI, GIÁM ĐỐC, NHÂN VIÊN, ĐỐI TÁC HOẶC
              NHÀ CUNG CẤP CỦA CHÚNG TÔI SẼ KHÔNG CHỊU TRÁCH NHIỆM VỀ BẤT KỲ
              THIỆT HẠI TRỰC TIẾP, GIÁN TIẾP, NGẪU NHIÊN, ĐẶC BIỆT, HẬU QUẢ HOẶC
              MANG TÍNH TRỪNG PHẠT, BAO GỒM NHƯNG KHÔNG GIỚI HẠN Ở:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Mất lợi nhuận, dữ liệu, sử dụng hoặc lợi thế vô hình khác</li>
              <li>Chi phí mua hàng hóa hoặc dịch vụ thay thế</li>
              <li>Truy cập trái phép hoặc thay đổi dữ liệu của bạn</li>
              <li>Tuyên bố hoặc hành vi của bên thứ ba</li>
            </ul>
            <p className="text-muted-foreground mt-4">
              GIỚI HẠN NÀY ÁP DỤNG BẤT KỂ CHÚNG TÔI CÓ ĐƯỢC THÔNG BÁO VỀ KHẢ
              NĂNG XẢY RA THIỆT HẠI ĐÓ HAY KHÔNG.
            </p>
          </CardContent>
        </Card>

        {/* Indemnification */}
        <Card>
          <CardHeader>
            <CardTitle>10. Bồi Thường</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Bạn đồng ý bảo vệ, bồi thường và giữ cho chúng tôi và các công ty
              liên kết, giám đốc, nhân viên, đại lý, đối tác và người được cấp
              phép không bị tổn hại khỏi và chống lại bất kỳ khiếu nại, trách
              nhiệm pháp lý, thiệt hại, tổn thất và chi phí nào, bao gồm phí
              luật sư hợp lý, phát sinh từ hoặc liên quan đến:
            </p>
            <ul className="list-disc list-inside space-y-1 text-muted-foreground ml-4">
              <li>Việc sử dụng dịch vụ của bạn</li>
              <li>Vi phạm các Điều Khoản này</li>
              <li>Vi phạm quyền của bên thứ ba</li>
              <li>Nội dung bạn đăng tải</li>
            </ul>
          </CardContent>
        </Card>

        {/* Governing Law */}
        <Card>
          <CardHeader>
            <CardTitle>11. Luật Điều Chỉnh</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Các Điều Khoản này sẽ được điều chỉnh và giải thích theo luật pháp
              Việt Nam. Bất kỳ tranh chấp nào phát sinh từ hoặc liên quan đến
              các Điều Khoản này sẽ được giải quyết tại tòa án có thẩm quyền tại
              Việt Nam.
            </p>
          </CardContent>
        </Card>

        {/* Dispute Resolution */}
        <Card>
          <CardHeader>
            <CardTitle>12. Giải Quyết Tranh Chấp</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Trong trường hợp có bất kỳ tranh chấp nào, chúng tôi khuyến khích
              bạn liên hệ với chúng tôi trước để tìm cách giải quyết thân thiện.
              Nếu không thể giải quyết, tranh chấp sẽ được giải quyết thông qua
              trọng tài hoặc tòa án có thẩm quyền.
            </p>
          </CardContent>
        </Card>

        {/* Severability */}
        <Card>
          <CardHeader>
            <CardTitle>13. Tính Độc Lập</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Nếu bất kỳ điều khoản nào trong các Điều Khoản này được cho là
              không hợp lệ hoặc không thể thi hành, điều khoản đó sẽ được giới
              hạn hoặc loại bỏ ở mức tối thiểu cần thiết và các điều khoản còn
              lại sẽ vẫn có hiệu lực đầy đủ.
            </p>
          </CardContent>
        </Card>

        {/* Waiver */}
        <Card>
          <CardHeader>
            <CardTitle>14. Từ Bỏ Quyền</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Việc chúng tôi không thực hiện hoặc thi hành bất kỳ quyền hoặc
              điều khoản nào trong các Điều Khoản này sẽ không cấu thành sự từ
              bỏ quyền hoặc điều khoản đó. Bất kỳ sự từ bỏ nào cũng chỉ có hiệu
              lực nếu được thực hiện bằng văn bản và được ký bởi đại diện được
              ủy quyền của chúng tôi.
            </p>
          </CardContent>
        </Card>

        {/* Changes to Terms */}
        <Card>
          <CardHeader>
            <CardTitle>15. Thay Đổi Điều Khoản</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Chúng tôi có quyền sửa đổi hoặc thay thế các Điều Khoản này bất kỳ
              lúc nào. Chúng tôi sẽ thông báo cho bạn về các thay đổi quan trọng
              ít nhất 30 ngày trước khi các điều khoản mới có hiệu lực bằng cách
              đăng thông báo trên dịch vụ hoặc gửi email. Việc bạn tiếp tục sử
              dụng dịch vụ sau khi các thay đổi có hiệu lực đồng nghĩa với việc
              bạn chấp nhận các Điều Khoản mới.
            </p>
          </CardContent>
        </Card>

        {/* Entire Agreement */}
        <Card>
          <CardHeader>
            <CardTitle>16. Toàn Bộ Thỏa Thuận</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Các Điều Khoản này, cùng với Chính Sách Bảo Mật và bất kỳ chính
              sách nào khác được đăng trên dịch vụ, cấu thành toàn bộ thỏa thuận
              giữa bạn và chúng tôi liên quan đến việc sử dụng dịch vụ và thay
              thế mọi thỏa thuận trước đó.
            </p>
          </CardContent>
        </Card>

        {/* Contact */}
        <Card>
          <CardHeader>
            <CardTitle>17. Liên Hệ</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-muted-foreground">
              Nếu bạn có bất kỳ câu hỏi hoặc thắc mắc nào về các Điều Khoản này,
              vui lòng liên hệ với chúng tôi qua:
            </p>
            <div className="space-y-2 text-muted-foreground">
              <p>
                <strong>Email:</strong> support@niensu.viet
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
            Bằng việc sử dụng dịch vụ Niên Sử Việt, bạn xác nhận rằng bạn đã đọc
            và hiểu các Điều Khoản Sử Dụng này và đồng ý bị ràng buộc bởi chúng.
          </p>
        </div>
      </div>
    </div>
  );
}
