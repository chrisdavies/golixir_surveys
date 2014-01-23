defmodule SurveyWriter do
  def writecsv({sections, questions}) do
    CSVLixir.write(transform(sections ++ questions))
  end

  defp transform([head | tail], result // []) do
    transform(tail, result ++ [to_array(head)])
  end

  defp transform([], result) do
    result 
  end

  defp to_array(Section[title: title, id: id]) do
    ["section", title, "", Kernel.integer_to_binary(id)]
  end

  defp to_array(Question[title: title, type: type, section: id]) do
    ["question", title, type, Kernel.integer_to_binary(id)]
  end

end