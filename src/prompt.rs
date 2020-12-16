use std::io::{self, Write};

pub fn prompt() -> Option<String> {
  let mut answer = String::new();

  match io::stdin().read_line(&mut answer) {
    Ok(0) => None,
    Ok(_) => Some(answer),
    Err(err) => {
      println!("error: {}", err);

      None
    }
  }
}

pub fn prompt_with_question(question: &str) -> Option<String> {
  print!("{}? ", question);

  if let Err(error) = std::io::stdout().flush() {
    println!("error: {}", error);
  }

  prompt()
}