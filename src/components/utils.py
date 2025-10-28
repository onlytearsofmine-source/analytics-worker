# utils.py
import logging
import os
import json
import requests

logger = logging.getLogger(__name__)

def get_config() -> dict:
    try:
        with open('config.json') as config_file:
            return json.load(config_file)
    except FileNotFoundError:
        logger.error('Config file not found')
        raise SystemExit(1)

def send_request(url: str, data: dict) -> requests.Response:
    try:
        response = requests.post(url, json=data)
        response.raise_for_status()
        return response
    except requests.exceptions.RequestException as e:
        logger.error(f'Request error: {e}')
        raise

def get_env_var(name: str) -> str:
    return os.environ.get(name)

def load_env_file(env_file: str) -> dict:
    try:
        with open(env_file) as env_file_content:
            return {line.strip().split('=')[0]: line.strip().split('=')[1] for line in env_file_content.readlines()}
    except FileNotFoundError:
        logger.error('Environment file not found')
        raise SystemExit(1)