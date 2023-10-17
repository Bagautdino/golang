import tkinter as tk
from tkinter import filedialog, ttk
import os
import platform
import subprocess
import winreg
import psutil
from cryptography.fernet import Fernet
import shutil

def install():
    # Сбор информации о системе
    user = os.getlogin()
    computer_name = os.environ['COMPUTERNAME']
    os_version = platform.platform()
    virtual_memory = psutil.virtual_memory()
    memory = virtual_memory.total / (1024. ** 3)  # Общий объем RAM в ГБ
    cpu = platform.processor()

# Сформированные данные для записи
    data = f"User: {user}\nComputer Name: {computer_name}\nOS Version: {os_version}\nMemory: {memory}GB\nCPU: {cpu}"

# Генерация ключа для шифрования
    key = Fernet.generate_key()
    cipher_suite = Fernet(key)

# Шифрование данных
    encrypted_data = cipher_suite.encrypt(data.encode())
    folder = filedialog.askdirectory(title="Выберите папку для установки")

# Запись зашифрованных данных в sys.tat
    with open(f'{folder}/sys.tat', 'wb') as file:
        file.write(encrypted_data)
        
    shutil.copy('secur.exe', folder)

# Сохранение ключа в реестре (или в другом месте по вашему усмотрению)
    registry_key = winreg.CreateKey(winreg.HKEY_CURRENT_USER, r"Software\Bagautdinov")
    winreg.SetValueEx(registry_key, "EncryptionKey", 0, winreg.REG_BINARY, key)
    # Закрыть инсталлятор
    root.destroy()

root = tk.Tk()
root.title("Installer")

label = tk.Label(root, text="Установка обновления")
label.pack(pady=10)

progress = ttk.Progressbar(root, orient="horizontal", length=300, mode="determinate")
progress.pack(pady=20)
progress.start()

install_button = tk.Button(root, text="Установить", command=install)
install_button.pack(pady=20)

root.mainloop()
