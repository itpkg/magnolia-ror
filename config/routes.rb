Rails.application.routes.draw do


  resources :leave_words, except: [:show, :edit, :update]
  resources :notices, except: :show

  # Dashboard
  get 'dashboard' => 'dashboard#index'
  %w(logs status users).each { |act| get "dashboard/#{act}" }
  %w(info seo nav_bar).each do |act|
    get "dashboard/#{act}"
    post "dashboard/#{act}"
  end
  delete 'dashboard/cache'

  # rate
  post 'rate' => 'home#rate'

  # search engine
  get 'google(*id).html', to: 'home#google'
  get 'baidu_verify_(*id).html', to: 'home#baidu'


  # Engines
  devise_for :users

  require 'sidekiq/web'
  authenticate :user, lambda { |u| u.is_admin? } do
    mount Sidekiq::Web => '/jobs'
  end

  # home
  root 'home#index'

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
