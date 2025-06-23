from pathlib import Path
import sys

CURRENT_DIR = Path(__file__).resolve().parent
sys.path.insert(0, str(CURRENT_DIR / "post_gen"))

# Import runner.main from inside post_gen
from runner import main

if __name__ == "__main__":
    main()
