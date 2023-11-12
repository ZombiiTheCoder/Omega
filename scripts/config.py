import os

def getJDK():
    JDK = 'E:/Program Files/Java/jdk-17/bin'
    if not os.path.exists(JDK):
        JDK = 'C:/Program Files/Java/jdk-17/bin'
    return JDK
