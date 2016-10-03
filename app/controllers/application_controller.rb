class ApplicationController < ActionController::Base
  include Pundit

  protect_from_forgery with: :exception

  # locale
  before_action :set_locale

  def set_locale
    # I18n.locale = params[:locale] || request.env['HTTP_ACCEPT_LANGUAGE'].split(',').first || I18n.default_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end

  def default_url_options
    {locale: I18n.locale}
  end

end
