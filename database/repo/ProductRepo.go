package repo

import (
	"fmt"

	"github.com/TechMaster/golang/15GoMySQL/model"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func initProduct() {
	// var sony, xiaomi model.Manufacturer
	// Db.Where("name = ?", "Sony").First(&sony)
	// Db.Where("name = ?", "Xiaomi").First(&xiaomi)

	var shoesSportMen, shoesSportGirl, shoesMen, shoesGirl model.Category
	Db.Where("name LIKE ?", "Giày thể thao nam%").First(&shoesSportMen)
	Db.Where("name LIKE ?", "Giày thể thao nữ%").First(&shoesSportGirl)
	Db.Where("name LIKE ?", "Giày da nam%").First(&shoesMen)
	Db.Where("name LIKE ?", "Giày da nữ%").First(&shoesGirl)

	shoes_1 := model.Product{
		Name: "Giày nam [FREESHIP EXTRA] Giày sneaker nam cổ cao kiểu dáng thể thao",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)
		`,
		Image:          "public/img/TTM1.jpg",
		PriceOld:       420000,
		PriceCurrent:   255000,
		Sale:           39,
		Madein:         "VN",
		CategoryID:     1,
		Category:       &shoesSportMen,
		HistoricalSold: 1300,
		RatingStar:     5,
		Quantity:       1,
	}

	shoes_2 := model.Product{
		Name: "Giày thể thao nam Giá rẻ Mác R8 Đế Siêu Êm",
		Description: `
		•	Thông tin sản phẩm: DOE SHOP
		•	Kiểu dáng: thể thao nam hàng nhập
		•	Size: 38-39-40-41-42-43
		•	Màu: đa dạng
		•	Xuất xứ: TQ
		•	Kiểu dáng trẻ trung năng động
		•	Bảo hành 1 đổi 1 trong vòng 1 tháng
		•	Dễ dàng phối đồ, dễ dàng vệ sinh, 
		•	Shop giầy sam sam cam kết bán sản phẩm chuẩn 
		•	Tới tay người mua hàng
		•	Chuẩn ảnh chuẩn video 100% do shoop tạo
		•	Uy tín tạo chất lượng, tạo thương hiệu`,
		Image:          "public/img/TTM2.jpg",
		PriceOld:       200000,
		PriceCurrent:   160000,
		Sale:           20,
		Madein:         "VN",
		CategoryID:     1,
		Category:       &shoesSportMen,
		HistoricalSold: 938,
		RatingStar:     4,
		Quantity:       1,
	}

	shoes_3 := model.Product{
		Name: "Giày da nam, thời trang cao cấp phong cách Hàn Quốc( ảnh thật)",
		Description: `
		INBOX VỚI SHOP ĐỂ TƯ VẤN SIZE)
		•	Giới thiệu về Giày Da Nam Trang Cao Cấp - GDB26
		•	Với thiết kế mũi giày tròn, đế được làm từ cao su tổng hợp cao cấp hạn chế mòn đế, thiết kế giúp các chàng trai cải thiện từ 3-5cm chiều cao. 
		Giày Da nam cao cấp được làm từ chất liệu da PU cao cấp cùng với lớp lót da mềm bên trong không mang đến cảm giác đau chân khi phải di chuyển nhiều. 
		•	Dù đi làm, đi học, tham gia sự kiện sang trọng hay chỉ đơn giản là dạo phố thì đôi giày da này luôn là sự lựa chọn đơn giản cho bạn bởi sự kết hợp nào cũng mang đến cho bạn sự chỉn chu, 
		lịch lãm nhưng vẫn vô cùng trẻ trung và cá tính. Quá xứng đáng để em nó được có mặt trong tủ giày của bạn đúng không nào.
		•	2, Tư vấn chọn size
		•	Giày Da có các Size: 38->44
		•	Inb hoặc liên hệ với shop để được tư vấn chọn size miễn phí.
		•	#GiàyDa  #GiàyDaCaoCấp #GiàyCôngSở  #GiàySneakerDa #GiàyĐộnĐế #giayfdanam #giayfdabong "`,
		Image:          "public/img/TTM3.jpg",
		PriceOld:       500000,
		PriceCurrent:   350000,
		Sale:           30,
		Madein:         "VN",
		CategoryID:     1,
		Category:       &shoesSportMen,
		HistoricalSold: 1938,
		RatingStar:     5,
		Quantity:       1,
	}

	shoes_4 := model.Product{
		Name: "Giày nam Nike XM355",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTM4.jpg",
		PriceOld:       200000,
		PriceCurrent:   105000,
		Sale:           50,
		Madein:         "VN",
		CategoryID:     1,
		Category:       &shoesSportMen,
		HistoricalSold: 938,
		RatingStar:     4,
		Quantity:       1,
	}

	shoes_5 := model.Product{
		Name: "giày thể thao nam - BHK123 ",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTM5.jpg",
		PriceOld:       189000,
		PriceCurrent:   172000,
		Sale:           13,
		Madein:         "US",
		CategoryID:     1,
		Category:       &shoesSportMen,
		HistoricalSold: 1138,
		RatingStar:     4,
		Quantity:       1,
	}

	TTG_1 := model.Product{
		Name: "Giày Sneaker Nữ MHMS02 Giày Thể Thao Nữ Chạy Bộ Siêu Thoáng Siêu Êm",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTG1.jpg",
		PriceOld:       300000,
		PriceCurrent:   169000,
		Sale:           44,
		Madein:         "US",
		CategoryID:     2,
		Category:       &shoesSportGirl,
		HistoricalSold: 1218,
		RatingStar:     5,
		Quantity:       1,
	}

	TTG_2 := model.Product{
		Name: "Giày Thể Thao Nữ Nữ MHMX1",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTG2.jpg",
		PriceOld:       189000,
		PriceCurrent:   17900,
		Sale:           5,
		Madein:         "VN",
		CategoryID:     2,
		Category:       &shoesSportGirl,
		HistoricalSold: 818,
		RatingStar:     4,
		Quantity:       1,
	}

	TTG_3 := model.Product{
		Name: "Giày Thể Thao Nữ ADG12",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTG3.jpg",
		PriceOld:       150000,
		PriceCurrent:   115000,
		Sale:           23,
		Madein:         "JP",
		CategoryID:     2,
		Category:       &shoesSportGirl,
		HistoricalSold: 2118,
		RatingStar:     5,
		Quantity:       1,
	}

	TTG_4 := model.Product{
		Name: "Giày Thể Thao Nữ POA65",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTG4.jpg",
		PriceOld:       430000,
		PriceCurrent:   215000,
		Sale:           50,
		Madein:         "JP",
		CategoryID:     2,
		Category:       &shoesSportGirl,
		HistoricalSold: 589,
		RatingStar:     5,
		Quantity:       1,
	}

	TTG_5 := model.Product{
		Name: "Giày Thể Thao Nữ YYYZ",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/TTG5.jpg",
		PriceOld:       225000,
		PriceCurrent:   179000,
		Sale:           20,
		Madein:         "VN",
		CategoryID:     2,
		Category:       &shoesSportGirl,
		HistoricalSold: 589,
		RatingStar:     4,
		Quantity:       1,
	}

	GDM1 := model.Product{
		Name: "Giày da nam công sở cao cấp màu đen buộc dây",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDM1.jpg",
		PriceOld:       420000,
		PriceCurrent:   380000,
		Sale:           10,
		Madein:         "VN",
		CategoryID:     3,
		Category:       &shoesMen,
		HistoricalSold: 1008,
		RatingStar:     5,
		Quantity:       1,
	}

	GDM2 := model.Product{
		Name: "Giày Da nam TTTTTTX",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDM2.jpg",
		PriceOld:       390000,
		PriceCurrent:   295000,
		Sale:           24,
		Madein:         "VN",
		CategoryID:     3,
		Category:       &shoesMen,
		HistoricalSold: 818,
		RatingStar:     5,
		Quantity:       1,
	}

	GDM_3 := model.Product{
		Name: "Giày Da Nam KKKKKK1",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDM3.jpg",
		PriceOld:       550000,
		PriceCurrent:   329000,
		Sale:           40,
		Madein:         "JP",
		CategoryID:     3,
		Category:       &shoesMen,
		HistoricalSold: 3118,
		RatingStar:     5,
		Quantity:       1,
	}

	GDM_4 := model.Product{
		Name: "Giày Da Nam Hot",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDM4.jpg",
		PriceOld:       840000,
		PriceCurrent:   579000,
		Sale:           31,
		Madein:         "US",
		CategoryID:     3,
		Category:       &shoesMen,
		HistoricalSold: 1589,
		RatingStar:     4,
		Quantity:       1,
	}

	GDM_5 := model.Product{
		Name: "Giày Da Nam xxxxxxZ",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDM5.jpg",
		PriceOld:       380000,
		PriceCurrent:   359000,
		Sale:           8,
		Madein:         "VN",
		CategoryID:     3,
		Category:       &shoesMen,
		HistoricalSold: 589,
		RatingStar:     5,
		Quantity:       1,
	}

	GDG_1 := model.Product{
		Name: "Giày da nữ cao cổ khóa sau cao cấp-boot mũi vuông",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDG1.jpg",
		PriceOld:       420000,
		PriceCurrent:   380000,
		Sale:           10,
		Madein:         "VN",
		CategoryID:     4,
		Category:       &shoesGirl,
		HistoricalSold: 1008,
		RatingStar:     5,
		Quantity:       1,
	}

	GDG_2 := model.Product{
		Name: "Giày Da nữ A12s",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDG2.jpg",
		PriceOld:       300000,
		PriceCurrent:   150000,
		Sale:           50,
		Madein:         "VN",
		CategoryID:     4,
		Category:       &shoesGirl,
		HistoricalSold: 618,
		RatingStar:     4,
		Quantity:       1,
	}

	GDG_3 := model.Product{
		Name: "Giày Da Nữ M001",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDG3.jpg",
		PriceOld:       400000,
		PriceCurrent:   300000,
		Sale:           25,
		Madein:         "VN",
		CategoryID:     4,
		Category:       &shoesGirl,
		HistoricalSold: 1118,
		RatingStar:     5,
		Quantity:       1,
	}

	GDG_4 := model.Product{
		Name: "Giày Da Nữ Hot",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDG4.jpg",
		PriceOld:       600000,
		PriceCurrent:   300000,
		Sale:           50,
		Madein:         "VN",
		CategoryID:     4,
		Category:       &shoesGirl,
		HistoricalSold: 1889,
		RatingStar:     5,
		Quantity:       1,
	}

	GDG_5 := model.Product{
		Name: "Giày Da Nam xxxxxxZ",
		Description: `
		Thông tin sản phẩm
		✔️Thương hiệu: NO BRAND
		✔️Mã sản phẩm: JD03, JD04
		✔️Kích thước: 39 40 41 42 43 44
		✔️Phong cách: Giày thể thao nam cao cổ, thời trang thích hợp đi chơi, đi học
		✔️Màu sắc: Giày nam JD bên mình có phối màu đen trắng và xám xanh trắng
		✔️Mùa thích hợp: đã là thời trang thì phá tan thời tiết 4 mùa mua nào đi cũng hợp ạ
		✔️Với phong cách trẻ trung năng động sản phẩm Giày nam JD các bạn có thể phối với các loại quần: quần bò, quần lửng, quần kaki, quần jogger rất đẹp và ngầu ạ
		✔️Chất liệu làm nên đôi giày là da PU thoáng không bí chân và vải được dệt may chuẩn xác và được shop cắt chỉ thừa và đóng gói cẩn thận ạ
		✔️Đối tượng áp dụng: 1 – 100 tuổi, đam mê là bất tận ạ
		✔️Đế giày thể thao nam: làm từ chất liệu cao su non, êm, đàn hồi, bên trong có lót trợ lực rất êm ái và giúp hạn chế mồ hôi tốt.
		✔️ Freeship/hỗ trợ phí vận chuyển (khách không biết lấy mã miễn phí vận chuyển / inbox cho shop để được hướng dẫn)"`,
		Image:          "public/img/GDG5.jpg",
		PriceOld:       380000,
		PriceCurrent:   359000,
		Sale:           8,
		Madein:         "VN",
		CategoryID:     4,
		Category:       &shoesGirl,
		HistoricalSold: 589,
		RatingStar:     5,
		Quantity:       1,
	}

	Db.Create(&shoes_1)
	Db.Create(&GDG_5)
	Db.Create(&shoes_2)
	Db.Create(&GDM_5)
	Db.Create(&TTG_1)
	Db.Create(&GDM1)
	Db.Create(&shoes_4)
	Db.Create(&TTG_3)
	Db.Create(&GDG_4)
	Db.Create(&shoes_3)
	Db.Create(&GDM2)
	Db.Create(&GDG_1)
	Db.Create(&TTG_5)
	Db.Create(&TTG_4)
	Db.Create(&shoes_5)
	Db.Create(&GDG_2)
	Db.Create(&GDM_4)
	Db.Create(&TTG_2)
	Db.Create(&GDM_3)
	Db.Create(&GDG_3)
}

func FindProductById(Id int) *model.Product {
	var product *model.Product
	Db.Preload("Category").First(&product, Id)
	fmt.Println(product)
	return product
}
