module ApplicationHelper

  def top_links
    links = Setting.get_site_info :top_links
    if links.nil?
      return {}
    end
    links.split("\r\n").reduce({}) do |h, l|
      args = l.split(' ')
      if args.size == 2
        h[args.first.to_sym] = args.last
      end
      h
    end
  end

end
