use console::style;

/// Format the input ``list`` as a list of identifier. They aren't escaped.
///
/// If ``list`` is empty, return an empty String
pub fn format_str_id_list<S: AsRef<str>>(list: &[S]) -> String {
    list.iter()
        .map(|s| format!("{}", style(s.as_ref()).bold()))
        .collect::<Vec<String>>()
        .join(", ")
}

#[cfg(test)]
mod tests {
    use crate::display::list::format_str_id_list;
    use console::set_colors_enabled;

    #[test]
    fn test_list_format() {
        set_colors_enabled(false);

        let empty_list: &[String] = &[];
        assert_eq!(&format_str_id_list(&empty_list), "");
        assert_eq!(&format_str_id_list(&["hello".to_string()]), "hello");
        assert_eq!(&format_str_id_list(&["hello", "world"]), "hello, world")
    }
}
