# utils.py

import os
import logging
from typing import List, Dict
from datetime import datetime

class Utils:
    def __init__(self):
        self.logger = logging.getLogger(__name__)

    def get_current_date(self) -> str:
        return datetime.now().strftime("%Y-%m-%d")

    def get_current_time(self) -> str:
        return datetime.now().strftime("%H:%M:%S")

    def get_config(self) -> Dict:
        with open(os.path.join(os.path.dirname(__file__), 'config.json')) as f:
            return json.load(f)

    def get_data_files(self) -> List:
        return [f for f in os.listdir('.') if f.endswith('.csv')]

    def create_dir(self, path: str) -> None:
        try:
            os.makedirs(path)
            self.logger.info(f"Directory created: {path}")
        except FileExistsError:
            self.logger.info(f"Directory already exists: {path}")

    def get_parent_dir(self, path: str) -> str:
        return os.path.dirname(path)

    def get_file_name(self, path: str) -> str:
        return os.path.basename(path)