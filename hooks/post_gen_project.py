#!/usr/bin/env python3
"""
Tiny Cookiecutter post-generation hook.

If the reusable package `cookiecutter-postgen` (containing post_gen.main)
is not installed yet, install it with pip, then run main().
"""
import importlib
import subprocess
import sys
from pathlib import Path

PKG_SPEC = "haraka==0.2.6"

def ensure_installed() -> None:
    try:
        importlib.import_module("haraka")
    except ImportError:
        subprocess.check_call([sys.executable, "-m", "pip", "install", PKG_SPEC])

ensure_installed()

from haraka import main, PostGenConfig
if __name__ == "__main__":

    cfg = PostGenConfig(
        language     = "{{ cookiecutter.language }}",
        project_slug = "{{ cookiecutter.project_slug }}",
        author       = "{{ cookiecutter.author }}",
        description  = "{{ cookiecutter.description }}",
        project_dir  = Path.cwd(),
        swagger      = False
    )

    main(cfg)
