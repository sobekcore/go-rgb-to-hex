# RGB to Hex
RGB to Hex is a simple little program made with Go. The GUI is made with Fyne.io. The program takes 3 values from user (red, green, blue) and turns it into hex color value. It was created for easier front-end development without need to google it all the time. Super lightweight and small.

## How to use?
Simply take out the shortcut file in the main directory wherever you want and use it! If you'd like to change something locally run:
~~~
go build -ldflags="-H windowsgui" -o ./cmd/hex-to-rgb.exe
~~~
and to create production build run:
~~~
fyne install -icon ./assets/icon.png
~~~

## Contact and contribution
If you'd like to contribute in any way, feel free to! There are probably a lot of usefull features and fixes to add. Also if you have and question hit me up anywhere, i will respond as fast as i can.
