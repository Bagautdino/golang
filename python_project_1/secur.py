import winreg
import sys
import tkinter as tk
from cryptography.fernet import Fernet

def get_signature_and_key():
    try:
        registry_key = winreg.OpenKey(winreg.HKEY_CURRENT_USER, r"Software\Bagautdinov")
        signature, _ = winreg.QueryValueEx(registry_key, "Signature")
        encryption_key, _ = winreg.QueryValueEx(registry_key, "EncryptionKey")
        winreg.CloseKey(registry_key)
        return signature, encryption_key
    except:
        return None, None

def display_decrypted_data():
    signature, encryption_key = get_signature_and_key()

    if signature != "MY_SIGNATURE":
        print("Неверная подпись!")
        sys.exit()

    # Дешифрование данных
    cipher_suite = Fernet(encryption_key)
    with open('sys.tat', 'rb') as file:
        encrypted_data = file.read()

    try:
        decrypted_data = cipher_suite.decrypt(encrypted_data).decode()
        
        # Создание GUI для отображения дешифрованных данных
        root = tk.Tk()
        root.title("Decrypted Data")

        text_widget = tk.Text(root, wrap=tk.WORD, width=80, height=20)
        text_widget.insert(tk.END, decrypted_data)
        text_widget.pack(padx=10, pady=10)

        root.mainloop()
        
    except Exception as e:
        print(f"Ошибка дешифрования! {str(e)}")

if __name__ == "__main__":
    display_decrypted_data()
