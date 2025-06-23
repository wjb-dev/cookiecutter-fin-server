import sys
from pathlib import Path

# This file lives in hooks/, so add that to sys.path
CURRENT_DIR = Path(__file__).resolve().parent
sys.path.insert(0, str(CURRENT_DIR))

from post_gen import main

if __name__ == "__main__":
    main()
