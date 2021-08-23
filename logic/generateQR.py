import qrcode
import pay_by_square
import sys

# THIS IS USING LIBRARY OF https://github.com/matusf/pay-by-square
# MIT LICENSE MODIFIED

code = pay_by_square.generate(
    amount=float(sys.argv[1]),
    iban=(sys.argv[2]),
    swift=(sys.argv[3]),
    beneficiary_name=(sys.argv[4])
)

print(code)
img = qrcode.make(code)
img.save("payCodeQR.png")
# img.show()
