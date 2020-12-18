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
