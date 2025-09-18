#![windows_subsystem = "windows"]

mod assets;

use std::io::Cursor;

use rand::Rng;
use rand::seq::IndexedRandom;
use rodio::OutputStreamBuilder;

slint::include_modules!();

fn main() {
    let app = MainWindow::new().unwrap();

    let mut rng = rand::rng();

    let stream_handle = OutputStreamBuilder::open_default_stream().unwrap();
 
    app.on_play_meow(move || {
        let sound_path = assets::SOUNDS.choose(&mut rng).unwrap();
        let sound_file = Cursor::new(assets::SoundAssets::get(sound_path).unwrap().data);
        let sink = rodio::play(&stream_handle.mixer(), sound_file).unwrap();
        sink.set_speed(rng.random_range(0.8..=1.2));
        sink.detach();
    });

    app.run().unwrap();
}
