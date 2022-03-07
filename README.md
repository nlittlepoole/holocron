# Horcrux
[![forthebadge](https://forthebadge.com/images/badges/contains-technical-debt.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/uses-badges.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)


![](https://media.giphy.com/media/LLxwPAjfpLak8/giphy.gif)
## About
This is a tool data URIs representing key guardians, meant to embed brain keys in physical objects, called a Horcrux. Horcruxes are small programs that reveal a `treasure` when the right response to a prompt is provided. Horcrux programs are small enough to be embedded in QR codes (and therefore [physical objects](https://qalo.com/collections/qr-dog-id-tags)). 

The resulting Horcrux should not be posted on the public internet. Brain key's aren't considered secure for [blockchain wallets](https://en.bitcoin.it/wiki/Brainwallet). However, in this case you can create a brain key (Horcrux) that does not generate a wallet directly but rather points to a more secure wallet. On top of that, byy embedding the Horcrux in an object and keeping it off the public internet, we create a second factor an attacker has to breach in order to access the treasure. The attacker has to both know the password to the Horcrux but must also physically possess the Horcrux in order to steal the treasure.

# Quickstart

## Required Dependencies

- [Docker](https://docs.docker.com/get-docker/)


## Testing A Horcrux 

![test](https://i.ibb.co/RzgwVyg/horcrux-example-big.png | width=500)

- Download or take a picture of the QR code with your phone
- Use a [ZXing](https://zxing.org/w/decode.jspx) based decoder to decode the 
- Paste the resulting data uri in your browser 
- Save the page as an html file to your local machine
- Open the HTML file in your browser
- Type the correct response into `Response` and click `Reveal` to see secret

**Note: You must save the base 64 HTML as an html file after loading because the webcrypto APIs are locked from use during normal base64 execution**

## Creating a Horcrux

### Asumptions

- The information you'd like to protect (`Treasure`) is relatively short (under 100 characters)
- The information does not need to be frequently accessed, this is better for cold storage
- The answer to your prompt is not trivially guessible 

### Generate Encrypted Key

- `make serve HORCRUX_PROMPT="<INSERT YOUR PROMPT>" HORCRUX_TREASURE="<INSERT TREASURE TO PROTECT>"`
- Open [http://127.0.0.1:8090/app](http://127.0.0.1:8090/app)
- Click Encrypt to Generate Encryption Key
- Copy this key to use in creating a Horcrux

### Generate Program Base64

- `make serve HORCRUX_PROMPT="<INSERT YOUR PROMPT>" HORCRUX_ENCRYPTED="<INSERT ENCRYPTED KEY FROM PREVIOUS STEP>"`
- `./build/app.html` will contain your horcrux program. Open it in a browser to test it works.
- `./build/app.txt` will contain a base64 version of the Horcrux you can use to turn into a QR code

### Encode to QR Code
I recommend [https://www.the-qrcode-generator.com/](https://www.the-qrcode-generator.com/) for QR code generation. The base64s are large enough that these programs end up being relatively big QR codes (bigger than level 30) and many basic apps can't read them.

### Decoding QR Code

Because many basic apps can't read QR codes above level 30, I recommend using one that is [Zxing](https://zxing.org/w/decode.jspx) based. I have succesfully tested decoding the original image, a variant generated by reducing the resolution by 50%, and a variant generated by taking a photo of the original QR code. 

### Physical Storage

QR codes are a great mechanism for physical storage since they can be manifested physically in practically any way an image can. Options include:

- Printing on a [Metal Plate](https://bayphoto.com/)
- [Lamination](https://www.fedex.com/en-us/office/binding-laminating-finishing-services.html?cmp=KNC-10000002-0-0-0-FXO-US-US-EN-AISFXO121510430&gclid=Cj0KCQiA95aRBhCsARIsAC2xvfxyFgrJqhUobH4TRA4CIT3g1DxGe2nC575DHcMcY8M7K1ZqGhmgXh4aAjAzEALw_wcB&gclsrc=aw.ds)
- [Ceramic Tiling](https://www.zazzle.com/qr+code+tiles)
- [Archival Focused Paper](https://www.futurepkg.com/best-paper)
- [Optical Storage](https://www.amazon.com/Best-Sellers-External-CD-DVD-Drives/zgbs/pc/1292121011)

Harddrives, SSDs, and flashdrives aren't great archival storage methods. Flash drives don't last that long and tape requires climate control to remain in good condition. A metal plate is small enough to fit in a safe or safety deposit box while also being able to survive adverse climates or even fires. 

### Types of Treasures to Store

- Blockchain Keys
- Master Key for Password Manager or 2 Factor Authenticator
- Coordinates to buried treasure if you have some
- etc


