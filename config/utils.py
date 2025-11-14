import logging
import json
from datetime import datetime, timedelta
from typing import Dict, List

logger = logging.getLogger(__name__)

def parse_config(config_str: str) -> Dict:
    try:
        return json.loads(config_str)
    except json.JSONDecodeError as e:
        logger.error(f"Failed to parse config: {e}")
        return {}

def get_time_window(start_time: datetime, end_time: datetime) -> List[datetime]:
    time_window = []
    while start_time <= end_time:
        time_window.append(start_time)
        start_time += timedelta(hours=1)
    return time_window

def filter_data(data: List[Dict], key: str, value: str) -> List[Dict]:
    return [item for item in data if item.get(key) == value]

def merge_dicts(dict1: Dict, dict2: Dict) -> Dict:
    return {**dict1, **dict2}

def is_empty(data: List) -> bool:
    return len(data) == 0

def get_logger() -> logging.Logger:
    return logger

class AnalyticsWorkerException(Exception):
    pass

class ConfigException(AnalyticsWorkerException):
    def __init__(self, message: str):
        self.message = message
        super().__init__(message)

class DataException(AnalyticsWorkerException):
    def __init__(self, message: str):
        self.message = message
        super().__init__(message)