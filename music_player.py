import discord
from discord.ext import commands
from discord.ext.commands import bot
from bot import MoodBot

from youtube_music import YouTubeMusic


class MusicPlayer:
    @commands.command(name="play", help="To play song")
    async def play(ctx, url):
        try:
            if not MoodBot.is_connected_to_voice(ctx):
                print("joining to voice")
                await MoodBot.join(ctx)

            server = ctx.guild
            voice_channel: discord.VoiceClient = discord.utils.get(
                bot.voice_clients, guild=server
            )

            async with ctx.typing():
                filename = await YouTubeMusic.get_music_filename_by_url(
                    url, loop=bot.loop
                )
                print("NEW FILENAME", filename)

                voice_channel.play(
                    discord.FFmpegPCMAudio(executable="ffmpeg", source=filename)
                )
            await ctx.send("**Now playing:** {}".format(filename))
        except Exception as e:
            await ctx.send("The bot is not connected to a voice channel." + str(e))

    @commands.command(name="pause", help="This command pauses the song")
    async def pause(ctx):
        voice_client = ctx.message.guild.voice_client
        if voice_client.is_playing():
            await voice_client.pause()
        else:
            await ctx.send("The bot is not playing anything at the moment.")

    @commands.command(name="resume", help="Resumes the song")
    async def resume(ctx):
        voice_client = ctx.message.guild.voice_client
        if voice_client.is_paused():
            await voice_client.resume()
        else:
            await ctx.send(
                "The bot was not playing anything before this. Use play command"
            )

    @commands.command(name="stop", help="Stops the song")
    async def stop(ctx):
        voice_client = ctx.message.guild.voice_client
        if voice_client.is_playing():
            await voice_client.stop()
        else:
            await ctx.send("The bot is not playing anything at the moment.")
