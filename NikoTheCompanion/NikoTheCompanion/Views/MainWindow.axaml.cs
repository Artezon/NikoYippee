using Avalonia.Controls;
using Avalonia.Media.Imaging;
using Avalonia.Platform;
using NAudio.Wave;
using System;
using System.Collections.Generic;
using System.IO;
using System.Reflection;

namespace NikoTheCompanion.Views;

public partial class MainWindow : Window
{
    private const string audioResName = "NikoTheCompanion.Assets.yippee.mp3";
    private MemoryStream audioMemoryStream;
    private readonly List<(WaveOutEvent output, Mp3FileReader reader)> activePlayers = new();

    private Bitmap nikoNormal = new Bitmap(AssetLoader.Open(new Uri("avares://NikoTheCompanion/Assets/niko.png")));
    private Bitmap nikoHappy = new Bitmap(AssetLoader.Open(new Uri("avares://NikoTheCompanion/Assets/niko_pancakes.png")));

    public MainWindow()
    {
        InitializeComponent();

        InitializeAudio();

        YippeeButton.Click += (_, _) => PlayYippee();
    }

    private void InitializeAudio()
    {
        var assembly = Assembly.GetExecutingAssembly();
        using var embeddedStream = assembly.GetManifestResourceStream(audioResName);
        if (embeddedStream == null) return;

        audioMemoryStream = new MemoryStream();
        embeddedStream.CopyTo(audioMemoryStream);
    }

    private void PlayYippee()
    {
        if (activePlayers.Count > 50) return;

        var playbackStream = new MemoryStream(audioMemoryStream.ToArray());

        var reader = new Mp3FileReader(playbackStream);
        var outputDevice = new WaveOutEvent();
        outputDevice.Init(reader);
        outputDevice.Play();

        if (activePlayers.Count == 0)
        {
            Niko.Source = nikoHappy;
        }
        activePlayers.Add((outputDevice, reader));

        outputDevice.PlaybackStopped += (s, e) =>
        {
            outputDevice.Dispose();
            reader.Dispose();
            playbackStream.Dispose();
            activePlayers.RemoveAll(x => x.output == outputDevice);

            if (activePlayers.Count == 0)
            {
                Niko.Source = nikoNormal;
            }
        };
    }

    private void StopAll()
    {
        foreach (var (output, reader) in activePlayers)
        {
            output.Stop();
            output.Dispose();
            reader.Dispose();
        }
        activePlayers.Clear();
    }
}
