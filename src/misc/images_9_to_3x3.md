# 图片


**把9张图片3X3在一张A4纸上打印**  
```
apt-get install imagemagick
montage -geometry +15+15 -tile 3x3 pic.* foo.jpg
/*
-geometry +15+15 x, y 方向上的间隔15,单位应该是像素
-tile 3x3 平铺成3x3的样子
*/
```

