defexception CSVError, [description: nil] do
  def message(exception) do
    exception.description
  end
end
