#ifndef __COMDLLEXPORT_H__
#define __COMDLLEXPORT_H__
                     
#define Success         0               /* Function execute OK */
#define	Cancel			255				/* Function execute Cansle */
#define Failure         255             /* Function execute Fail */
#define OutTime         255             /* So long without echo, may be hardware lost */

#ifdef __cplusplus
extern "C" 
{
#endif

/*-----------------------------------------------------------
（1）   Name:           OnUSBHIDConnect
        Function:       打开USBHID接口
        Input:          
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall OnUSBHIDConnect(); 
/*-----------------------------------------------------------
（2）   Name:           OnUSBHIDDisconnect
        Function:       关闭USBHID接口
        Input:          
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
void __stdcall OnUSBHIDDisconnect(); 
//-----------------------------------------------------
//	USB HID PINPAD
//---------------------------------------------------
/*-----------------------------------------------------------
（3）   Name:           PINPAD_ACTIVE_USBHID
        Function:       激活工作密钥
        Input:          (int) MKeyNO: 主密钥号（0 - 32）
			(int) UKeyNO: 工作密钥号（0 - 7）
			(int) TSDESflag:加密标志：‘1’为3DES标志，其余为DES标志
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_ACTIVE_USBHID(int MKeyNO, int UKeyNO, int TSDESflag);
/*-----------------------------------------------------------
（4）   Name:           PINPAD_F_USBHID
        Function:       激活密码键盘
        Input:          (int) flab: '0'-请输入密码；'1'-请再输入密码
			(int*) outlen: 返回的密文长度
			(unsigned char *) output:返回的密文
			(int) Out_Time: 超时时间，单位：秒
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_F_USBHID(int flab, int *outlen, unsigned char *output, int Out_Time);
/*-----------------------------------------------------------
（5）   Name:           PINPAD_E_USBHID
        Function:       激活密码键盘
        Input:          (int) flab: '0'-请输入密码；'1'-请再输入密码
			(int) Out_Time: 超时时间，单位：秒
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_E_USBHID(int flab, int Out_Time);
/*-----------------------------------------------------------
（6）   Name:           PINPAD_G_USBHID
        Function:       获取密文
        Input:          (int*) outlen: 返回的密文长度
			(unsigned char *) output:返回的密文
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_G_USBHID(int *outlen, unsigned char *output);
/*-----------------------------------------------------------
（7）   Name:           PINPAD_H_USBHID
        Function:       字符串加密
        Input:          (int) flab: '0'-通过G命令获取密文；'1'-直接回送密文
			(unsigned char *) input:8字节待加密的字符串
			(unsigned char *) output:返回的密文
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_H_USBHID(int flab, unsigned char *input, unsigned char *output);
/*-----------------------------------------------------------
（8）   Name:           PINPAD_I_USBHID
        Function:       卡号参与加密
        Input:          (int) flab: '0'-请输入密码；'1'-请再输入密码；'2'-无语音提示
			(unsigned char *) CardNO:12字节的字符串（卡号）
			(int*) outlen:返回的密文长度
			(unsigned char *) output:返回的密文
			(int) Out_Time:超时时间，单位：秒
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_I_USBHID(int flab, unsigned char *CardNO, int *outlen, unsigned char *output, int Out_Time);
/*-----------------------------------------------------------
（9）   Name:           PINPAD_L_USBHID
        Function:       设置最长密码长度
        Input:          (int) len : 0<len<13
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_L_USBHID(int len);
/*-----------------------------------------------------------
（10）   Name:           PINPAD_M_USBHID
        Function:       修改主密钥
        Input:          (int) MKeyNO: 主密钥号
			(int) Keylen: 主密钥长度（8-DES,16-3DES）
			(unsigned char *) OldMKey:旧主密钥
			(unsigned char *) NewMKey:新主密钥
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_M_USBHID(int MKeyNO, int Keylen, unsigned char *OldMKey, unsigned char *NewMKey);
/*-----------------------------------------------------------
（11）   Name:           PINPAD_N_USBHID
        Function:       设置最短密码长度
        Input:          (int) len : 0<len<13
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_N_USBHID(int len);
/*-----------------------------------------------------------
（12）   Name:           PINPAD_R_USBHID
        Function:       设置最短密码长度
        Input:          (int) flab : 0-全部，1-指定主密钥区
			(int) KeyNO: 主密钥号（在flab不为0时有效）
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_R_USBHID(int flab, int KeyNO);
/*-----------------------------------------------------------
（13）   Name:           PINPAD_S_USBHID
        Function:       修改工作密钥
        Input:          (int) MKeyNO: 主密钥号
			(int) NKeyNO: 工作密钥号
			(int) Strlen: 字符串长度（8-DES,16-3DES）
			(unsigned char *) Str:字符串
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_S_USBHID(int MKeyNO, int NKeyNO, int Strlen, unsigned char *Str);
/*-----------------------------------------------------------
（14）   Name:           PINPAD_T_USBHID
        Function:       计算MAC
        Input:          (int) flab: 1-ECB算法;0-X99算法
			(int) Strlen: 字符串长度
			(unsigned char *) Str:字符串
			(int *) outlen:MAC长度
			(unsigned char *) output:MAC
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_T_USBHID(int flab, int Strlen, unsigned char *Str, int *outlen, unsigned char *output);
/*-----------------------------------------------------------
（15）   Name:           PINPAD_V_USBHID
        Function:       获取版本信息
        Input:          (unsigned char *) Str:版本信息
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_V_USBHID(unsigned char *Str);
/*-----------------------------------------------------------
（16）   Name:           GET_PINPAD
		Function:       获取版本信息
		Input:          int PinLen:密码长度
		Input:          char * Str:版本信息
		Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall GET_PINPAD(int PinLen,char *Str);
/*-----------------------------------------------------------
（17）   Name:           GET_PINPAD_
		Function:       获取版本信息
		Input:          int PinLen:密码长度
		Input:          char * Str:版本信息
		Input:          int flab:声音提示'0'-请输入密码；'1'-请再输入密码
		Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall GET_PINPAD_(int PinLen, char * Str, int flab);
#ifdef __cplusplus
}
#endif

#endif __COMDLLEXPORT_H__