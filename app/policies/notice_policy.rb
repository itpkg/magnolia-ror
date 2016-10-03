class LeaveWordPolicy < ApplicationPolicy
    def create?
      !user.nil? && user.is_admin?
    end

    def update?
      !user.nil? && user.is_admin?
    end

    def destroy?
      !user.nil? && user.is_admin?
    end
  end
