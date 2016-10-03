# RailsSettings Model
class Setting < RailsSettings::Base
  source Rails.root.join('config/app.yml')
  namespace Rails.env

  def self.get_site_info(key)
    Setting["#{I18n.locale}://site//#{key}"] || nil
  end

  def self.set_site_info(key, val)
    Setting["#{I18n.locale}://site//#{key}"] = val
  end

end
