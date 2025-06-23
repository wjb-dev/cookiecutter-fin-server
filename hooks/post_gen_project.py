from pathlib import Path
import sys

sys.path.insert(0, str(Path.cwd() / "hooks"))

from post_gen import main

if __name__ == "__main__":
    main()
