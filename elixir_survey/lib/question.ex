defrecord Question, title: nil, type: nil, section: 0

defmodule QuestionType do
	def range, do: 0
	def para, do: 1

	def from_string("para") do
		para
	end

	def from_string("1-5") do
		range
	end

	def to_string(0) do
		"1-5"
	end

	def to_string(1) do
		"para"
	end
end