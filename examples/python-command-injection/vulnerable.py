import subprocess
import sys

def ping_host(host):
    subprocess.run("ping -c 1 " + host, shell=True)

if __name__ == "__main__":
    if len(sys.argv) > 1:
        ping_host(sys.argv[1])
