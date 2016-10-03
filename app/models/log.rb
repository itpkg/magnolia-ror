class Log < ApplicationRecord
  belongs_to :user
  validates :message, presence: true

  enum flag: [info: 0, notice: 1, error: 2]
end
