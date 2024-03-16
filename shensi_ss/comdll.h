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
��1��   Name:           OnUSBHIDConnect
        Function:       ��USBHID�ӿ�
        Input:          
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall OnUSBHIDConnect(); 
/*-----------------------------------------------------------
��2��   Name:           OnUSBHIDDisconnect
        Function:       �ر�USBHID�ӿ�
        Input:          
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
void __stdcall OnUSBHIDDisconnect(); 
//-----------------------------------------------------
//	USB HID PINPAD
//---------------------------------------------------
/*-----------------------------------------------------------
��3��   Name:           PINPAD_ACTIVE_USBHID
        Function:       �������Կ
        Input:          (int) MKeyNO: ����Կ�ţ�0 - 32��
			(int) UKeyNO: ������Կ�ţ�0 - 7��
			(int) TSDESflag:���ܱ�־����1��Ϊ3DES��־������ΪDES��־
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_ACTIVE_USBHID(int MKeyNO, int UKeyNO, int TSDESflag);
/*-----------------------------------------------------------
��4��   Name:           PINPAD_F_USBHID
        Function:       �����������
        Input:          (int) flab: '0'-���������룻'1'-������������
			(int*) outlen: ���ص����ĳ���
			(unsigned char *) output:���ص�����
			(int) Out_Time: ��ʱʱ�䣬��λ����
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_F_USBHID(int flab, int *outlen, unsigned char *output, int Out_Time);
/*-----------------------------------------------------------
��5��   Name:           PINPAD_E_USBHID
        Function:       �����������
        Input:          (int) flab: '0'-���������룻'1'-������������
			(int) Out_Time: ��ʱʱ�䣬��λ����
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_E_USBHID(int flab, int Out_Time);
/*-----------------------------------------------------------
��6��   Name:           PINPAD_G_USBHID
        Function:       ��ȡ����
        Input:          (int*) outlen: ���ص����ĳ���
			(unsigned char *) output:���ص�����
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_G_USBHID(int *outlen, unsigned char *output);
/*-----------------------------------------------------------
��7��   Name:           PINPAD_H_USBHID
        Function:       �ַ�������
        Input:          (int) flab: '0'-ͨ��G�����ȡ���ģ�'1'-ֱ�ӻ�������
			(unsigned char *) input:8�ֽڴ����ܵ��ַ���
			(unsigned char *) output:���ص�����
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_H_USBHID(int flab, unsigned char *input, unsigned char *output);
/*-----------------------------------------------------------
��8��   Name:           PINPAD_I_USBHID
        Function:       ���Ų������
        Input:          (int) flab: '0'-���������룻'1'-�����������룻'2'-��������ʾ
			(unsigned char *) CardNO:12�ֽڵ��ַ��������ţ�
			(int*) outlen:���ص����ĳ���
			(unsigned char *) output:���ص�����
			(int) Out_Time:��ʱʱ�䣬��λ����
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_I_USBHID(int flab, unsigned char *CardNO, int *outlen, unsigned char *output, int Out_Time);
/*-----------------------------------------------------------
��9��   Name:           PINPAD_L_USBHID
        Function:       ��������볤��
        Input:          (int) len : 0<len<13
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_L_USBHID(int len);
/*-----------------------------------------------------------
��10��   Name:           PINPAD_M_USBHID
        Function:       �޸�����Կ
        Input:          (int) MKeyNO: ����Կ��
			(int) Keylen: ����Կ���ȣ�8-DES,16-3DES��
			(unsigned char *) OldMKey:������Կ
			(unsigned char *) NewMKey:������Կ
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_M_USBHID(int MKeyNO, int Keylen, unsigned char *OldMKey, unsigned char *NewMKey);
/*-----------------------------------------------------------
��11��   Name:           PINPAD_N_USBHID
        Function:       ����������볤��
        Input:          (int) len : 0<len<13
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_N_USBHID(int len);
/*-----------------------------------------------------------
��12��   Name:           PINPAD_R_USBHID
        Function:       ����������볤��
        Input:          (int) flab : 0-ȫ����1-ָ������Կ��
			(int) KeyNO: ����Կ�ţ���flab��Ϊ0ʱ��Ч��
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_R_USBHID(int flab, int KeyNO);
/*-----------------------------------------------------------
��13��   Name:           PINPAD_S_USBHID
        Function:       �޸Ĺ�����Կ
        Input:          (int) MKeyNO: ����Կ��
			(int) NKeyNO: ������Կ��
			(int) Strlen: �ַ������ȣ�8-DES,16-3DES��
			(unsigned char *) Str:�ַ���
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_S_USBHID(int MKeyNO, int NKeyNO, int Strlen, unsigned char *Str);
/*-----------------------------------------------------------
��14��   Name:           PINPAD_T_USBHID
        Function:       ����MAC
        Input:          (int) flab: 1-ECB�㷨;0-X99�㷨
			(int) Strlen: �ַ�������
			(unsigned char *) Str:�ַ���
			(int *) outlen:MAC����
			(unsigned char *) output:MAC
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_T_USBHID(int flab, int Strlen, unsigned char *Str, int *outlen, unsigned char *output);
/*-----------------------------------------------------------
��15��   Name:           PINPAD_V_USBHID
        Function:       ��ȡ�汾��Ϣ
        Input:          (unsigned char *) Str:�汾��Ϣ
        Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall PINPAD_V_USBHID(unsigned char *Str);
/*-----------------------------------------------------------
��16��   Name:           GET_PINPAD
		Function:       ��ȡ�汾��Ϣ
		Input:          int PinLen:���볤��
		Input:          char * Str:�汾��Ϣ
		Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall GET_PINPAD(int PinLen,char *Str);
/*-----------------------------------------------------------
��17��   Name:           GET_PINPAD_
		Function:       ��ȡ�汾��Ϣ
		Input:          int PinLen:���볤��
		Input:          char * Str:�汾��Ϣ
		Input:          int flab:������ʾ'0'-���������룻'1'-������������
		Output:         Success , Failure and OutTime
--------------------------------------------------------------*/
int __stdcall GET_PINPAD_(int PinLen, char * Str, int flab);
#ifdef __cplusplus
}
#endif

#endif __COMDLLEXPORT_H__