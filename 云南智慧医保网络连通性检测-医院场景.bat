@echo off
color 2F
title �����ǻ�ҽ��������ͨ�Լ��-ҽԺ����
rem yangwl neusoft
echo.
echo ��ʼ���. . .
echo.
echo.
ping -n 2 10.114.24.201>%temp%\1.ping
findstr "TTL" %temp%\1.ping>nul
if %errorlevel%==0 (echo     �� ҽ��ר������) else (echo     �� ҽ��ר��������������ҽԺ����Ӫ�����磡)
echo.
ping -n 2 10.11.200.205>%temp%\2.ping
findstr "TTL" %temp%\2.ping>nul
if %errorlevel%==0 (echo     �� ���Ҿ�DNS������10.11.200.205��������) else (echo     �� ���Ҿ�DNS������10.11.200.205���ʲ�����������ҽԺ����Ӫ�����磡)
echo.
ping -n 2 10.116.130.21>%temp%\3.ping
findstr "TTL" %temp%\3.ping>nul
if %errorlevel%==0 (echo     �� ���Ͼ�DNS������10.116.130.21��������) else (echo     �� ���Ͼ�DNS������10.116.130.21���ʲ�����������ҽԺ����Ӫ�����磡)
echo.
ping -n 2 ldjk.yn.hsip.gov.cn>%temp%\4.ping
findstr "TTL" %temp%\4.ping>nul
if %errorlevel%==0 (echo     �� �����ӿ�����ldjk.yn.hsip.gov.cn��������) else (echo     �� �����ӿ�����ldjk.yn.hsip.gov.cn���ʲ�����������DNS���ã�)
echo.
ping -n 2 ec.yn.hsip.gov.cn>%temp%\5.ping
findstr "TTL" %temp%\5.ping>nul
if %errorlevel%==0 (echo     �� ����ƾ֤����ec.yn.hsip.gov.cn��������) else (echo     �� ����ƾ֤����ec.yn.hsip.gov.cn���ʲ�����������DNS���ã�)
echo.
ping -n 2 ld.yn.hsip.gov.cn>%temp%\6.ping
findstr "TTL" %temp%\6.ping>nul
if %errorlevel%==0 (echo     �� �����Ż�����ld.yn.hsip.gov.cn��������) else (echo     �� �����Ż�����ld.yn.hsip.gov.cn���ʲ�����������DNS���ã�)
echo.
ping -n 2 10.114.161.37>%temp%\7.ping
findstr "TTL" %temp%\7.ping>nul
if %errorlevel%==0 (echo     �� ������������֤��ַ��������) else (echo     �� ������������֤��ַ���ʲ�����������ҽԺ����Ӫ�����磡)
echo.
if exist %temp%\*.ping del %temp%\*.ping
echo.
echo.
pause