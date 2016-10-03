class LeaveWordsController < ApplicationController

  def new
    @leave_word = LeaveWord.new
    render 'form'
  end

  def create
    @leave_word = LeaveWord.create params.require(:leave_word).permit(:content)
    if @leave_word.valid?
      flash[:notice] = ' '
    else
      flash[:alert] = @leave_word.errors
    end
    redirect_to new_leave_word_path
  end

  def destroy
    lw = LeaveWord.find params[:id]
    authorize lw
    lw.destroy

    redirect_to leave_words_path
  end

  def index
    @items = LeaveWord.order(id: :desc).page params[:page]
    authorize @items
  end

end
