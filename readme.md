#### 本服务用于解析并301跳转srv dns纪录。
#### 用法:
  - 浏览器打开 /srv/$(SRV_RECORD)  =>  301 跳转到 http://$(SRV_RECORD)
  - 浏览器打开 /srv/$(SRV_RECORD)?https=true  => 301 跳转到 https://$(SRV_RECORD)

#### This service is used to parse and 301 redirect srv DNS records.
#### Usage: 
  - Open /srv/$(SRV_RECORD) in a browser => 301 redirect to http://$(SRV_RECORD)
  - Open /srv/$(SRV_RECORD)?https=true in a browser => 301 redirect to https://$(SRV_RECORD)

#### about dns srv record：
  - https://zh.wikipedia.org/zh-hans/Template:SRV
  - https://baike.baidu.com/item/SRV%E8%AE%B0%E5%BD%95/10637211