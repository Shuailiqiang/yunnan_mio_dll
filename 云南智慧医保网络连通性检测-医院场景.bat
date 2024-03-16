@echo off
color 2F
title 云南智慧医保网络连通性检测-医院场景
rem yangwl neusoft
echo.
echo 开始检测. . .
echo.
echo.
ping -n 2 10.114.24.201>%temp%\1.ping
findstr "TTL" %temp%\1.ping>nul
if %errorlevel%==0 (echo     √ 医保专网正常) else (echo     × 医保专网不正常，需检查医院或运营商网络！)
echo.
ping -n 2 10.11.200.205>%temp%\2.ping
findstr "TTL" %temp%\2.ping>nul
if %errorlevel%==0 (echo     √ 国家局DNS服务器10.11.200.205访问正常) else (echo     × 国家局DNS服务器10.11.200.205访问不正常，需检查医院或运营商网络！)
echo.
ping -n 2 10.116.130.21>%temp%\3.ping
findstr "TTL" %temp%\3.ping>nul
if %errorlevel%==0 (echo     √ 云南局DNS服务器10.116.130.21访问正常) else (echo     × 云南局DNS服务器10.116.130.21访问不正常，需检查医院或运营商网络！)
echo.
ping -n 2 ldjk.yn.hsip.gov.cn>%temp%\4.ping
findstr "TTL" %temp%\4.ping>nul
if %errorlevel%==0 (echo     √ 两定接口域名ldjk.yn.hsip.gov.cn访问正常) else (echo     × 两定接口域名ldjk.yn.hsip.gov.cn访问不正常，需检查DNS配置！)
echo.
ping -n 2 ec.yn.hsip.gov.cn>%temp%\5.ping
findstr "TTL" %temp%\5.ping>nul
if %errorlevel%==0 (echo     √ 电子凭证域名ec.yn.hsip.gov.cn访问正常) else (echo     × 电子凭证域名ec.yn.hsip.gov.cn访问不正常，需检查DNS配置！)
echo.
ping -n 2 ld.yn.hsip.gov.cn>%temp%\6.ping
findstr "TTL" %temp%\6.ping>nul
if %errorlevel%==0 (echo     √ 两定门户域名ld.yn.hsip.gov.cn访问正常) else (echo     × 两定门户域名ld.yn.hsip.gov.cn访问不正常，需检查DNS配置！)
echo.
ping -n 2 10.114.161.37>%temp%\7.ping
findstr "TTL" %temp%\7.ping>nul
if %errorlevel%==0 (echo     √ 三代卡联机认证地址访问正常) else (echo     × 三代卡联机认证地址访问不正常，需检查医院或运营商网络！)
echo.
if exist %temp%\*.ping del %temp%\*.ping
echo.
echo.
pause