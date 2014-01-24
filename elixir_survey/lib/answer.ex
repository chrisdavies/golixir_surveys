defmodule AnswerValidator do
	def validate(questions, answers) do
		if length(questions) != length(answers) do 
			raise ValidationError, description: "The number of answers does not match the number of questions"
		end

		verify_answers questions, answers
	end

	defp verify_answers([Question[type: 0] |questions], [answer |answers]) do
		if (answer != nil && (answer < 1 || answer > 5)) do
			raise ValidationError, description: "Range answers must be between 1 and 5"
		end

		verify_answers questions, answers
	end

	defp verify_answers([Question[type: 1] |questions], [answer |answers]) do
		if answer != nil && !is_binary answer do
			raise ValidationError, description: "Para answers must be text"
		end

		verify_answers questions, answers
	end

	defp verify_answers([], []) do
	end
end