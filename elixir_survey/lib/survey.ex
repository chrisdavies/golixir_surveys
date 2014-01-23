defrecord Survey, 
	title: nil, 
	is_write_once: false, 
	is_anonymous: true, 
	active_after: nil, 
	active_until: nil,
	sections: [],
	questions: []

defmodule SurveyValidator do
	def validate(survey) do
		validate_required_fields(survey)
		validate_dates(survey)
		validate_questions(survey)
	end

	defp validate_required_fields(survey) do 
		not_nil(survey.title, "Title")
		not_nil(survey.active_until, "Active until")
		not_nil(survey.active_after, "Active after")
		not_nil(survey.sections, "Sections")
		not_nil(survey.questions, "Questions")
	end

	defp validate_dates(survey) do
		if survey.active_after > survey.active_until do
			raise ValidationError, description: "Active until must be later than active after"
		end
	end

	defp validate_questions(survey) do
		sectionIds = HashDict.new Enum.map(survey.sections, fn(s) -> {s.id, nil} end)
		badQuestion = Enum.find(survey.questions, fn(q) -> !Dict.has_key?(sectionIds, q.section) end)
		if badQuestion do 
			raise ValidationError, description: "Question '#{badQuestion.title}' references non-exisistant section #{badQuestion.section}"
		end
	end

	defp not_nil(val, field) do
		if val == nil do
			raise ValidationError, description: "#{field} is required"
		end
	end
end