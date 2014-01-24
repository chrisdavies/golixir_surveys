defmodule SurveyReader do
  def readcsv(records) do
  	process_line(CSVLixir.read(records), {[], []})
  end

  defp process_line([], records) do 
  	records
  end

  defp process_line([[] | []], records) do 
  	records
  end

  defp process_line([line | tail], records) do
  	case line do 
  		[rec_type, title, q_type, section_id] ->
  			process_line(tail, add_record(lower(rec_type), title, q_type, section_id, records))
  		_ -> raise CSVError, description: "Records must be in the format 'record-type, title, question-type, section-id'"
  	end
  end

  defp add_record("section", title, _, id, {sections, questions}) do
  	{[Section.new(title: strip(title), id: to_int(id))] ++ sections, questions}
  end

  defp add_record("question", title, type, section_id, {sections, questions}) do
  	type = lower(type)
    title = strip(title)

  	cond do
      title == "" -> raise CSVError, description: "Title cannot be blank"
  		type == "1-5" || type == "para" -> 
  			{sections, [Question.new(title: strip(title), type: type, section: to_int(section_id))] ++ questions}
  		true -> raise CSVError, description: "Unknown question type #{type}"
  	end
  end

  defp add_record(record_type, _, _, _, _) when record_type != "section" and record_type != "question" do
  	raise CSVError, description: "Unknown record type #{record_type}"
  end

  defp lower(str) do
  	String.downcase(strip(str))
  end

  defp strip(str) do
  	String.strip(str)
  end

  defp to_int(str) do
  	try do
  		Kernel.binary_to_integer(strip(str))
  	rescue
  		ArgumentError -> raise CSVError, description: "Section id '#{str}' must be an integer"
  	end
  end

end
