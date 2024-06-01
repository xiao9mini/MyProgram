import time
import os
# cd to the directory where the file is located

# get the directory where the file is located


directory = os.path.dirname(__file__)
os.chdir(directory)

# def rename_file(file_name, new_file_name):
file_name = "proxy"
s = os.listdir()
t = int(time.time())
file_pre2new = {}
for file_name in os.listdir():
    if file_name.endswith(".list"):
        file_pre = file_name.split(".")[0]
        new_file_name = f"{file_pre}.{t}.list"
        os.rename(file_name, new_file_name)
        file_pre2new[file_pre] = new_file_name

new_ini = []
with open("config.ini", "r") as f:
    for line in f:
        if "MyProgram" in line:
            file_end = line.split("/")[-1]
            file_pre = file_end.split(".")[0]
            new_file_name = file_pre2new[file_pre]
            new_line = line.replace(file_end, new_file_name)
            new_ini.append(new_line+ "\n")
        else:
            new_ini.append(line)

with open("config.ini", "w") as f:
    f.writelines(new_ini)
