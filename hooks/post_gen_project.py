from pathlib import Path
import sys

# Add the hooks directory to the sys.path
HOOKS_DIR = Path(__file__).resolve().parent
sys.path.insert(0, str(HOOKS_DIR))

from post_gen.runner import main

if __name__ == "__main__":
    main()