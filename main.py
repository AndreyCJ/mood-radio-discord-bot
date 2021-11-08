import os

from bot import MoodBot
from music_player import MusicPlayer


COMMAND_PREFIX = "!"
DISCORD_TOKEN = os.getenv("DISCORD_TOKEN")


if __name__ == "__main__":
    MoodBot(DISCORD_TOKEN, COMMAND_PREFIX)
    # MusicPlayer(MoodBot)
