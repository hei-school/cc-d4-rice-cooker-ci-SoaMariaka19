require 'minitest/autorun'
require_relative 'riceCooker'

class RiceCookerTest < Minitest::Test
  def setup
    @console_spy = []
    def @console_spy.log(msg)
      self << msg
    end

    def @console_spy.restore
      @spied_obj = nil
    end
  end

  def set_rice_and_water_added(value)
    RiceCooker.instance_variable_set(:@rice_and_water_added, value)
  end

  def test_initial_state
    assert_equal RiceCooker::STATE_OFF, RiceCooker.instance_variable_get(:@rice_cooker_state)
    refute RiceCooker.instance_variable_get(:@rice_and_water_added)
  end

  def test_display_state
    RiceCooker.display_state
    assert_includes @console_spy, "\nCurrent state of the rice cooker: #{RiceCooker::STATE_OFF}"
  end

  def test_plug_in_without_rice_and_water
    set_rice_and_water_added(false)
    RiceCooker.plug_in
    assert_includes @console_spy, 'Error: Add rice and water before plugging in and starting cooking.'
    refute RiceCooker.instance_variable_get(:@rice_cooker_state) == RiceCooker::STATE_COOKING
  end

  def test_plug_in
    set_rice_and_water_added(true)
    RiceCooker.plug_in
    assert_includes @console_spy, 'The rice cooker is plugged in, and cooking begins.'
    assert_equal RiceCooker::STATE_COOKING, RiceCooker.instance_variable_get(:@rice_cooker_state)
  end

  def test_plug_in_already_cooking
    set_rice_and_water_added(true)
    RiceCooker.instance_variable_set(:@rice_cooker_state, RiceCooker::STATE_COOKING)
    RiceCooker.plug_in
    assert_includes @console_spy, 'Error: The rice cooker is already cooking or in keep warm mode.'
    assert_equal RiceCooker::STATE_COOKING, RiceCooker.instance_variable_get(:@rice_cooker_state)
  end

  def test_plug_in_already_keep_warm
    set_rice_and_water_added(true)
    RiceCooker.instance_variable_set(:@rice_cooker_state, RiceCooker::STATE_KEEP_WARM)
    RiceCooker.plug_in
    assert_includes @console_spy, 'Error: The rice cooker is already cooking or in keep warm mode.'
    assert_equal RiceCooker::STATE_KEEP_WARM, RiceCooker.instance_variable_get(:@rice_cooker_state)
  end

  def test_finish_cooking
    set_rice_and_water_added(true)
    RiceCooker.instance_variable_set(:@rice_cooker_state, RiceCooker::STATE_COOKING)
    RiceCooker.finish_cooking
    assert_includes @console_spy, 'Cooking is finished. The rice cooker is in keep warm mode.'
    assert_equal RiceCooker::STATE_KEEP_WARM, RiceCooker.instance_variable_get(:@rice_cooker_state)
  end

  def test_finish_cooking_no_cooking_in_progress
    set_rice_and_water_added(true)
    RiceCooker.finish_cooking
    assert_includes @console_spy, 'Error: No cooking in progress.'
    assert_equal RiceCooker::STATE_OFF, RiceCooker.instance_variable_get(:@rice_cooker_state)
  end

  def test_quit_program
    RiceCooker.quit_program
    assert_includes @console_spy, "\n======================================================================="
    assert_includes @console_spy, 'Goodbye! Thanks for using the rice cooker program.'
  end

  def teardown
    @console_spy.restore
  end
end
