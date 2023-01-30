import os
import sys

params = ["account_in", "project_in", "version_in"]


def get_params(params: list):
    for num in range(1, 4):
        try:
            params[num-1] = sys.argv[num]
        except:
            params[num-1] = None


def run_cmd(params: list):
    try:
        if None not in params:
            os.system(f"docker build -t ghcr.io/{params[0]}/{params[1]}:v{params[2]} .")
            os.system(f"docker push ghcr.io/{params[0]}/{params[1]}:v{params[2]}")
    except Exception as e:
        raise e


if __name__ == "__main__":
    get_params(params)
    run_cmd(params)
