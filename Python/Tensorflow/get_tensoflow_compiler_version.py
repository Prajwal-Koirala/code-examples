import tensorflow



# Get the current tensorflow complier version.
def get_tensoflow_compiler_version():
    return tensorflow.version.COMPILER_VERSION



def main():
    print(get_tensoflow_compiler_version())


    

main()
