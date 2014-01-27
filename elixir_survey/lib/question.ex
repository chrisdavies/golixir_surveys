defrecord Question, title: nil, type: nil, section: 0

defmodule QuestionType do
	def to_string(:para), do: "para"
	def to_string(:range), do: "1-5"

	def to_atom("para"), do: :para
	def to_atom("1-5"), do: :range
end