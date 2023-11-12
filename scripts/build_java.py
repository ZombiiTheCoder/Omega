import os, shutil, subprocess, zlib, zipfile

def copyDir(src, dst):
    try :
        shutil.copytree(src, dst)
        print('Directory copied successfully.')
    except Exception as e:
        print('Directory not copied.')
        print(e)

def removeRec(dir):
    for root, paths, files in os.walk(dir):
        for x in paths:
            removeRec(os.path.join(root, x))
        for x in files:
            os.remove(os.path.join(root, x))
    os.rmdir(dir)

def removeRecExt(dir, ext):
    for root, paths, files in os.walk(dir):
        for x in paths:
            removeRecExt(os.path.join(root, x), ext)
        for x in files:
            if ext in x:
                os.remove(os.path.join(root, x))

def getAllFiles(dir):
    f = set()
    fr = []
    for root, paths, files in os.walk(dir):
        for x in paths:
            fr.extend(getAllFiles(os.path.join(root, x)))
        for x in files:
            f.add(os.path.join(root, x))

    for x in fr:
        f.add(x)
    return f

def getAllFilesExt(dir, ext):
    f = set()
    fr = []
    for root, paths, files in os.walk(dir):
        for x in paths:
            f = getAllFilesExt(os.path.join(root, x), ext)
            fr.extend(f)
        for x in files:
            if ext in x:
                f.add(os.path.join(root, x))

    for x in fr:
        f.add(x)
    return f

def Compress(file, fileout):
    f = open(file, "rb")
    bin_string = f.read()
    f.close()
    compressed = zlib.compress(bin_string)
    f2 = open(fileout, "wb")
    f2.write(compressed)
    f2.close()

def Decompress(file, fileout):
    f = open(file, "rb")
    bina = f.read()
    f.close()
    compressed = zlib.decompress(bina)
    f2 = open(fileout, "wb")
    f2.write(compressed)
    f2.close()

def CompileWithoutLibs(JarName, src, srcBin):
    # JarName = "PrimalUtils"; src = "PrimalUtils"; srcBin = "PrimalUtilsBin"
    copyDir(src, srcBin)
    Src, _ = getAllFilesExt(srcBin, ".java")
    Builder = ["javac"]; Builder.extend(Src)
    subprocess.call(Builder)
    for x in Src: os.remove(x)
    os.chdir(srcBin)
    os.system("jar cf ../bin/"+JarName+".jar *")
    os.chdir("../"); removeRec(srcBin)
    Compress("bin/"+JarName+".jar", "bin/"+JarName+".bin")
    return "bin/"+JarName+".jar"

def CompileWithLibs(JarName, src, srcBin, LibsPath, LibsExtra = []):
    # JarName = "PrimalCraft-0.1a"; src = "src"; srcBin = "srcBin"
    copyDir(src, srcBin)
    Src = getAllFilesExt(srcBin, ".java")
    Libs = getAllFilesExt(LibsPath, ".jar")
    for x in LibsExtra: Libs.add(x)
    Builder = ["javac"]; Builder.extend(Src); Builder.append("-cp"); Builder.append(";".join(Libs))
    # print(" ".join(Builder))
    subprocess.call(Builder)
    for x in Src: os.remove(x)
    os.chdir(srcBin)
    os.system("jar cf ../bin/"+JarName+".jar *")
    os.chdir("../"); removeRec(srcBin)
    # Compress("bin/"+JarName+".jar", "bin/"+JarName+".bin")
    return "bin/"+JarName+".jar"

gv = "Omega"
PrimalCraft = CompileWithLibs(gv, "src", "srcBin", "lib")

if not os.path.exists("bin"): os.mkdir("bin")

Libs = getAllFilesExt("lib", ".jar")

# java -cp PrimalCraft-0.1a.jar org.geostudios.java.main.Main
# print("bin/"+gv+".jar;"+(";".join(Libs)))
subprocess.call(["java", "-cp",  "bin/"+gv+".jar;"+(";".join(Libs)), "org.zombii.java.main.Main"])
os.chdir("../")