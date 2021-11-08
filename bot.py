import json
import urllib.request
import discord
from discord.ext import commands
from discord.ext.commands import CommandNotFound


intents = discord.Intents().all()
client = discord.Client(intents=intents)


class MoodBot:
    def __init__(self, discord_token, command_prefix):
        self.discord_bot = commands.Bot(command_prefix=command_prefix, intents=intents)

        self.discord_bot.run(discord_token)

    async def say_random_phrase(ctx):
        try:
            headers = {
                "Content-Type": "application/json",
                "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_4) AppleWebKit/605.1.15 "
                "(KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
                "Origin": "https://yandex.ru",
                "Referer": "https://yandex.ru/",
            }

            API_URL = "https://zeapi.yandex.net/lab/api/yalm/text3"
            payload = {"query": ctx.message.content[1:], "intro": 0, "filter": 1}
            params = json.dumps(payload).encode("utf8")

            async with ctx.typing():
                phrase_req = urllib.request.Request(
                    API_URL, data=params, headers=headers
                )
                phrase_res = urllib.request.urlopen(phrase_req)
                phrase = json.loads(phrase_res.read().decode("utf8"))

                print("phrase =>", phrase)
                await ctx.message.channel.send(phrase["text"])
        except Exception as e:
            print("Cant get balabola", e)

    @commands.command(name="join", help="Tells the bot to join the voice channel")
    async def join(ctx):
        if not ctx.message.author.voice:
            await ctx.send(
                "{} is not connected to a voice channel".format(ctx.message.author.name)
            )
            return
        else:
            channel = ctx.message.author.voice.channel
        await channel.connect()

    @commands.command(name="leave", help="To make the bot leave the voice channel")
    async def leave(ctx):
        voice_client = ctx.message.guild.voice_client
        if voice_client.is_connected():
            await voice_client.disconnect()
        else:
            await ctx.send("The bot is not connected to a voice channel.")

    @commands.command(name="балабала", help="Says text")
    async def sayBalalba(self, ctx):
        await self.say_random_phrase(ctx)

    @commands.command(name="эй", help="Says hello")
    async def sayHello(ctx):
        await ctx.message.channel.send("Задрова нахуййййй!!!!!!111111")

    @client.event
    async def on_ready():
        print("We have logged in as {0.user}".format(client))

    @client.event
    async def on_command_error(self, ctx, error):
        if isinstance(error, CommandNotFound):
            await self.say_random_phrase(ctx)
            return

        raise error

    # @client.event
    # async def on_message(message):
    #     if message.author == client.user:
    #         return

    #     if message.content.startswith('$hello'):
    #         await message.channel.send('Hello!')

    def is_connected_to_voice(ctx):
        voice_client = discord.utils.get(ctx.discord_bot.voice_clients, guild=ctx.guild)
        return voice_client and voice_client.is_connected()
