use rust_embed::RustEmbed;

#[derive(RustEmbed)]
#[folder = "assets/sounds"]
pub struct SoundAssets;

#[derive(RustEmbed)]
#[folder = "assets/images"]
pub struct ImageAssets;

pub static SOUNDS: [&str; 1] = [
    "cat_3.ogg",
];
